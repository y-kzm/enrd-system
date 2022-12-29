package app

import (
	"context"
	"database/sql"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/vishvananda/netlink"
	"github.com/vishvananda/netlink/nl"

	"github.com/y-kzm/enrd-system/api"
	meas_client "github.com/y-kzm/enrd-system/pkg/tool/client"
)

const database = "enrd:0ta29SourC3@tcp(controller:3306)/enrd"

var (
	Nic   string
	Store []*api.SRInfo
)

type Server struct {
	api.UnimplementedServiceServer
}

// Recieve Configure message
func (s *Server) Configure(ctx context.Context, in *api.ConfigureRequest) (*api.ConfigureResponse, error) {
	log.Printf("Called configure procedure")
	log.Println(in.SrInfo) // debug
	if in.Msg == "go" {
		// TODO: テーブル名とパスの対応付けをしとく必要あり?>in.SrInfoを覚えとけばOK？Successのときだけ別の変数で記憶しとく？
		for i, _ := range in.SrInfo {
			if err := IPv6AddrAdd(in.SrInfo[i].SrcAddr, Nic); err != nil {
				// TODO: Cleanup()
				log.Print(err)
				return &api.ConfigureResponse{
					Status: 1,
					Msg:    "Failed to assign IPv6 address",
				}, nil
			}
			if err := CreateVRF(in.SrInfo[i].Vrf, in.SrInfo[i].SrcAddr); err != nil {
				// TODO: Cleanup()
				log.Print(err)
				return &api.ConfigureResponse{
					Status: 1,
					Msg:    "Failed to create VRF",
				}, nil
			}
			if err := SEG6EncapRouteAdd(in.SrInfo[i].DstAddr, in.SrInfo[i].Vrf, Nic, in.SrInfo[i].SidList); err != nil {
				// TODO: Cleanup()
				log.Print(err)
				return &api.ConfigureResponse{
					Status: 1,
					Msg:    "Failed to add seg6 encap route",
				}, nil
			}
		}
		Store = in.SrInfo
		return &api.ConfigureResponse{
			Status: 0,
			Msg:    "Success!!!",
		}, nil
	} else {
		return &api.ConfigureResponse{
			Status: 1,
			Msg:    "Bad Msg...",
		}, nil
	}
}

// Recieve Measure message
func (s *Server) Measure(ctx context.Context, in *api.MeasureRequest) (*api.MeasureResponse, error) {
	log.Printf("Called measure procedure")
	log.Print(Store) // debug
	if in.Method == "ptr" {
		// TODO: 各パスの測定を行う 測定パス数のループ
		srcIP, _, err := net.ParseCIDR(Store[0].SrcAddr)
		if err != nil {
			log.Print(err)
			return &api.MeasureResponse{
				Status: 1,
				Msg:    "Failed to start measurement",
			}, nil
		}
		dstIP, _, err := net.ParseCIDR(Store[0].DstAddr)
		if err != nil {
			log.Print(err)
			return &api.MeasureResponse{
				Status: 1,
				Msg:    "Failed to start measurement",
			}, nil
		}
		res := meas_client.EstimateClient(int(in.Param.PacketNum), int(in.Param.RepeatNum), int(in.Param.PacketSize), srcIP.String(), dstIP.String())
		log.Print(res)
		return &api.MeasureResponse{
			Status: 0,
			Msg:    "Success!!!",
		}, nil
	} else {
		return &api.MeasureResponse{
			Status: 1,
			Msg:    "Not supported",
		}, nil
	}
}

// Connection to database
func ConnectDB() error {
	db, err := sql.Open("mysql", database)
	if err != nil {
		return err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return err
	} else {
		var version string
		db.QueryRow("SELECT VERSION()").Scan(&version)
		log.Println("DB Version: ", version)
	}

	return nil
}

// TODO: CleanUp()
/**
 * Ex.
 *  ip -6 rule del from fd00:0:172:16:ffff::1/64
 *  ip -6 route del fc00:4::/64  encap seg6 mode encap segs fc00:3::,fc00:2:: dev enp6s0 table 102 metric 1024 pref medium
 *  ip addr del fd00:0:172:16:ffff::1/64 dev enp6s0
 *  ip addr del fc00:1::/64 dev lo
 *  ip route del fc00:1::/64
 */
// func CleanUp() {}

// Addition of Seg6 End route
/**
 * Ex.
 *  ip route add fc00:3::/128 encap seg6local action End dev net0
 */
func SEG6LocalRouteEndAdd(dst string, dev string) error {
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
func IPv6AddrAdd(src string, dev string) error {
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
func CreateVRF(vrf int32, src string) error {
	srcIP, srcIPnet, err := net.ParseCIDR(src)
	if err != nil {
		return err
	}

	srcNet := &net.IPNet{IP: srcIP, Mask: srcIPnet.Mask}

	_, err = netlink.RuleList(netlink.FAMILY_V6)
	if err != nil {
		return err
	}

	rule := netlink.NewRule()
	rule.Table = int(vrf)
	rule.Src = srcNet
	rule.Priority = 5
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
func SEG6EncapRouteAdd(dst string, vrf int32, dev string, sidlist []string) error {
	var sidList []net.IP

	seg6encap := &netlink.SEG6Encap{Mode: nl.SEG6_IPTUN_MODE_ENCAP}
	for _, sid := range sidlist {
		ip, _, err := net.ParseCIDR(sid)
		if err != nil {
			return err
		}
		sidList = append(sidList, ip)

	}
	seg6encap.Segments = sidList

	link, err := netlink.LinkByName(dev)
	if err != nil {
		return err
	}
	dstIP, dstIPnet, err := net.ParseCIDR(dst)
	if err != nil {
		return err
	}

	route := netlink.Route{
		LinkIndex: link.Attrs().Index,
		Dst: &net.IPNet{
			IP:   dstIP,
			Mask: dstIPnet.Mask,
		},
		Encap: seg6encap,
		Table: int(vrf),
	}
	_ = netlink.RouteDel(&route)

	return netlink.RouteAdd(&route)
}
