package app

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"
)

const database = "enrd:0ta29SourC3@tcp(controller:3306)/enrd"

// Connection test to DB
func ConnectToDB() {
	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		var version string
		db.QueryRow("SELECT VERSION()").Scan(&version)
		log.Println("DB Version: ", version)
	}
}

// Assignment of SID
func AssignSID(sid string) {
	lo, err := netlink.LinkByName("lo")
	if err != nil {
		log.Fatal(err)
	}
	addr, err := netlink.ParseAddr(sid)
	if err != nil {
		log.Fatal(err)
	}
	err = netlink.AddrAdd(lo, addr)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: Remove unwanted loopback addresses
// func RemoveSID()

// TODO: Delete routes to the specified SID
// func RouteDel()

// Add End route
func SEG6LocalRouteEndAdd(dst string, dev string) {
	li, err := netlink.LinkByName(dev)
	if err != nil {
		log.Fatal(err)
	}

	dstIP, dstIPnet, err := net.ParseCIDR(dst)
	if err != nil {
		log.Fatal(err)
	}

	var flags_end [nl.SEG6_LOCAL_MAX]bool
	flags_end[nl.SEG6_LOCAL_ACTION] = true
	e := &netlink.SEG6LocalEncap{
		Flags:  flags_end,
		Action: nl.SEG6_LOCAL_ACTION_END,
	}
	route := netlink.Route{
		LinkIndex: li.Attrs().Index,
		Dst: &net.IPNet{
			IP:   dstIP,
			Mask: dstIPnet.Mask,
		},
		Encap: e,
	}

	if err := netlink.RouteAdd(&route); err != nil {
		log.Fatal(err)
	}
}
