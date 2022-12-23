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

// Connection to database
func ConnectDB() {
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

// TODO: CleanUp()
	// TODO: Remove unwanted loopback addresses
	// func RemoveSID()

	// TODO: Delete routes to the specified SID
	// func RouteDel()


// Addition of Seg6 End route
/** 
 * Ex.
 *  ip route add fc00:3::/128 encap seg6local action End dev net0
 */
func SEG6LocalRouteEndAdd(dst string, dev string) (error) {
	li, err := netlink.LinkByName(dev)
	if err != nil {
		return err
	}
	dstIP, dstIPnet, err := net.ParseCIDR(dst)
	if err != nil {
		return err
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
		return err
	}

	return nil
}


// Addition of IPv6 address to a specified interface
/** 
 * Ex.
 *  ip addr add fd00:0:172:16:4::12/64 dev net0
 *  ip addr add fc00:3::/64 dev lo
 */
func IPv6AddrAdd(src string, dev string) (error){
	link, err := netlink.LinkByName(dev)
	if err != nil {
		return err
	}
	addr, err := netlink.ParseAddr(src)
	if err != nil {
		return err
	}
	err = netlink.AddrAdd(link, addr)
	if err != nil {
		return err
	}

	return nil
}

// Creation of VRF per each source address
/** 
 * Ex.
 *  ip -6 rule add from fd00:0:172:16:2::5 table 100
 */
func CreateVRF(vrf int, src string) (error) {
	skipUnlessRoot(t)
	defer setUpNetlinkTest(t)()

	srcNet, err := netlink.ParseAddr(src)
	if err != nil {
		return err
	}
	// dstNet := &net.IPNet{IP: net.IPv4(172, 16, 1, 1), Mask: net.CIDRMask(24, 32)}

	rulesBegin, err := RuleList(FAMILY_V6)
	if err != nil {
		return err
	}

	rule := NewRule()
	rule.Table = vrf
	rule.Src = srcNet
	// rule.Dst = dstNet
	rule.Priority = 5
	// rule.OifName = main.nic
	// rule.IifName = main.nic
	if err := RuleAdd(rule); err != nil {
		return err
	}

	return nil
}

// Addition of Seg6 Encap route
/** 
 * Ex.
 *  ip route add fc00:4::/64 encap seg6 mode encap segs fc00:2::,fc00:3:: dev net0 table 100
 */
func SEG6EncapRouteAdd(dst string, vrf int, nic string, sidlist []string)











