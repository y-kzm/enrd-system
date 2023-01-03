package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/y-kzm/enrd-system/api"
	"github.com/y-kzm/enrd-system/cmd/agent/app"
	meas_server "github.com/y-kzm/enrd-system/pkg/tool/server"
)

const port = 52000

func main() {
	// Argument check
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: ./agent [InterfaceName] [SID]\n")
		os.Exit(1)
	}
	app.Nic = os.Args[1]
	log.Printf("Interface: %s", os.Args[1])
	log.Printf("Prefix-SID: %s", os.Args[2])

	// Connecting to the database
	if err := app.ConnectDB(); err != nil {
		log.Print("Failed to connect to database")
		os.Exit(1)
	}

	// TODO: Cleanup()
	// 1. 不要なLoopbackアドレスを削除
	// 2. 指定SID宛のルートを削除

	// Assignment of SID
	if err := app.IPv6AddrAdd(os.Args[2], "lo"); err != nil {
		log.Print("Failed to assign SID")
		// TODO: Cleanup()
		os.Exit(1)
	}
	// Adding an End Route
	if err := app.SEG6LocalRouteEndAdd(os.Args[2], app.Nic); err != nil {
		log.Print("Failed to add End route")
		// TODO: Cleanup()
		os.Exit(1)
	}

	// Startup of IGI/PTR server
	// TODO: エラー処理
	go meas_server.EstimateServer()
	log.Print("Startup of IGI/PTR server")

	// Startup of gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Printf("Failed to listen: %v", err)
		// TODO: Cleanup()
		os.Exit(1)
	}
	s := grpc.NewServer()
	api.RegisterServiceServer(s, &app.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Printf("Failed to serve: %v", err)
		// TODO: Cleanup()
		os.Exit(1)
	}
}
