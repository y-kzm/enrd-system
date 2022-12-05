package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"

	"github.com/y-kzm/enrd-system/api"
	"github.com/y-kzm/enrd-system/cmd/agent/app"
	tool "github.com/y-kzm/enrd-system/pkg/tool/server"
)

const database = "root:0ta29SourC3@tcp(127.0.0.1:3306)/enrd"
const port = 52000

func main() {
	if len(os.Args) != 3 {
		log.Fatalf("Argument error: %#v", os.Args)
	}
	log.Println("NIC: %s SID: %s", os.Args[1], os.Args[2])

	// TODO: 指定のNICが存在するかどうかチェック
	// TODO: SIDがIPv6形式かどうかチェック

	// Connect to db
	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check connectivity to DB
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	log.Println("DB Version: ", version)

	// TODO: クリーン処理

	// TODO: SID付与，End設定

	// TODO: IGI/PTR serverの起動
	go tool.EstimateServer()
	log.Print("Start IGI/PTR server")

	// Start gRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	api.RegisterServiceServer(s, &app.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
