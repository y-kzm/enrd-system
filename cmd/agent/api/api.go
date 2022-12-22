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
	log.Printf("Called Configure()")
	log.Print(in.SrInfo[0])
	log.Print(in.SrInfo[1])
	log.Print(in.SrInfo[2])
	log.Print(in.SrInfo[3])
	log.Print(in.SrInfo[4])
	if in.Msg == "go" {
		return &api.ConfigureResponse{
			Status: 0,
			Msg:    "OK!!!",
		}, nil
	} else {
		return &api.ConfigureResponse{
			Status: 1,
			Msg:    "NG...",
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
