package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/y-kzm/enrd-system/api"
	procedure "github.com/y-kzm/enrd-system/cmd/agent/api"
	"github.com/y-kzm/enrd-system/cmd/agent/app"
	tool "github.com/y-kzm/enrd-system/pkg/tool/server"
)

const port = 52000

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: ./agent [NIC] [SID]\n")
		log.Fatalf("Argument error: %#v", os.Args)
	}
	log.Printf("NIC: %s SID: %s", os.Args[1], os.Args[2])

	// Connect to database
	app.ConnectToDB()

	// TODO: クリーン処理
	// 1. 不要なLoopbackアドレスを削除
	// 2. 指定SID宛のルートを削除

	// Assign SID, Add End route
	app.AssignSID(os.Args[2])
	log.Print("Successful assign SID")
	app.SEG6LocalRouteEndAdd(os.Args[2], os.Args[1])
	log.Print("Successful add End route")

	// Start IGI/PTR server
	go tool.EstimateServer()
	log.Print("Start IGI/PTR server")

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterServiceServer(s, &procedure.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
