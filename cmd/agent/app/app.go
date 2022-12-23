package app

import (
	"database/sql"
	"log"
	"net"
	"context"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"

	"github.com/y-kzm/enrd-system/api"
)

const database = "enrd:0ta29SourC3@tcp(controller:3306)/enrd"

type Server struct {
	api.UnimplementedServiceServer
}

// Recieve Configure message
func (s *Server) Configure(ctx context.Context, in *api.ConfigureRequest) (*api.ConfigureResponse, error) {
	log.Printf("Called configure procedure")
	// debug
	log.Print(in.SrInfo)
	if in.Msg == "go" {
		// TODO: テーブル名とパスの対応付けをしとく必要あり?>in.SrInfoを覚えとけばOK？Successのときだけ別の変数で記憶しとく？
		for i, _ := range in.SrInfo {
			if err := IPv6AddrAdd(in.SrInfo[i].SrcAddr, "enp6s0"); err != nil {
				// TODO: Cleanup()
				return &api.ConfigureResponse{
					Status: 1,
					Msg: "Failed to assign IPv6 address",
				}, err
			}
			if err := CreateVRF(in.SrInfo[i].Vrf, in.SrInfo[i].SrcAddr); err != nil {
				// TODO: Cleanup()
				return &api.ConfigureResponse{
					Status: 1,
					Msg: "Failed to create VRF",
				}, err				
			}
			if err := SEG6EncapRouteAdd(in.SrInfo[i].DstAddr, in.SrInfo[i].Vrf, "enp6s0", in.SrInfo[i].SidList); err != nil {
				// TODO: Cleanup()
				return &api.ConfigureResponse{
					Status: 1,
					Msg: "Failed to add seg6 encap route",
				}, err						
			}
		}
		return &api.ConfigureResponse{
			Status: 0,
			Msg: "Success!!!",
		}, nil
	} else {
		return &api.ConfigureResponse{
			Status: 1,
			Msg: "Bad Msg...",
		}, nil
	}
}

// Recieve Measure message
func (s *Server) Measure(ctx context.Context, in *api.MeasureRequest) (*api.MeasureResponse, error) {
	log.Printf("Called Measure()")
	if in.Method == "ptr" {
		return &api.MeasureResponse{
			Status: 0,
			Msg:    "OK!!!",
		}, nil
	} else {
		return &api.MeasureResponse{
			Status: 1,
			Msg:    "NG...",
		}, nil
	}
}


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
func CreateVRF(vrf int32, src string) (error) {
	//ip1 := net.ParseIP("fd56:6b58:db28:2913::")
	//ip2 := net.ParseIP("fde9:379f:3b35:6635::")

	srcIP, srcIPnet, err := net.ParseCIDR(src)
	if err != nil {
		return err
	}
	dstIP, dstIPnet, err := net.ParseCIDR("::/128")
	if err != nil {
		return err
	}

	srcNet := &net.IPNet{IP: srcIP, Mask: srcIPnet.Mask}
	dstNet := &net.IPNet{IP: dstIP, Mask: dstIPnet.Mask}

	_, err = netlink.RuleList(netlink.FAMILY_V6)
	if err != nil {
		return err
	}

	rule := netlink.NewRule()
	rule.Table = int(vrf)
	rule.Src = srcNet
	rule.Dst = dstNet
	rule.Priority = 5
	// rule.OifName = main.nic
	// rule.IifName = main.nic
	if err := netlink.RuleAdd(rule); err != nil {
		return err
	}

	return nil
}

// Addition of Seg6 Encap route
/** 
 * Ex.
 *  ip route add fc00:4::/64 encap seg6 mode encap segs fc00:2::,fc00:3:: dev net0 table 100
 */
func SEG6EncapRouteAdd(dst string, vrf int32, nic string, sidlist []string) (error) {
	return nil
}











