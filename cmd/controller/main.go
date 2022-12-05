package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
	"github.com/y-kzm/enrd-system/pkg/commands"
)

const (
	name        = "enrd"
	description = "Estimating available bandwidth using SRv6"
)

var (
	version = "0.0.1"
)

const addr = "localhost:52000"

func main() {
	if err := newApp().Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = name
	app.Version = version
	app.Usage = description
	app.Authors = []*cli.Author{
		{
			Name:  "yykzm",
			Email: "yokoo@v6.netsci.info.hiroshima-cu.ac.jp",
		},
	}
	app.Commands = commands.Commands
	return app
}

/*
func main() {
	argv := os.Args
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Measure(ctx, &api.MeasureRequest{
		Method: argv[1],
	})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}
	log.Printf("Received from server: %d %s", r.GetStatus(), r.GetMsg())
}
*/
