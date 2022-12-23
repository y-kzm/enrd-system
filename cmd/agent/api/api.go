package procedure

import (
	"context"
	"log"

	"github.com/y-kzm/enrd-system/api"
)

type Server struct {
	api.UnimplementedServiceServer
}

// Recieve Configure message
func (s *Server) Configure(ctx context.Context, in *api.ConfigureRequest) (*api.ConfigureResponse, error) {
	log.Printf("Called configure procedure")
	// log.Print(in.SrInfo)
	if in.Msg == "go" {
		// TODO: テーブル名とパスの対応付けをしとく必要あり?>in.SrInfoを覚えとけばOK？Successのときだけ別の変数で記憶しとく？
		for i := len(in.SrInfo) {
			if err := IPv6AddrADD(in.SrInfo[i].SrcAddr, main.nic); err != nil {
				// TODO: Cleanup()
				return &api.ConfigureResponse{
					Status: 1,
					Msg: "Failed to assign IPv6 address",
				}, err
			}
			if err = CreateVRF(in.SrInfo[i].Vrf, in.SrInfo[i].SrcAddr); err != nil {
				// TODO: Cleanup()
				return &api.ConfigureResponse{
					Status: 1,
					Msg: "Failed to create VRF",
				}, err				
			}
			if err = SEG6EncapRouteAdd(in.SrInfo[i].DstAddr, in.SrInfo[i].Vrf, main.nic, in.SrInfo[i].SidList); err != nil {
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
