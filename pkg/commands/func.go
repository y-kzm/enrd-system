package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func CmdInit(c *cli.Context) error {
	// f, err := os.Open("/etc/enrd/config/temprate.yaml")
	f, err := os.Open("./config/temprate.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	buf := make([]byte, 1024)
	for {
		n, err := f.Read(buf)
		if n == 0 || err != nil {
			break
		}
		fmt.Println(string(buf[:n]))
	}

	return nil
}

func CmdConf(c *cli.Context) error {
	// yamlファイル解析関数へ
	fmt.Fprint(os.Stdout, "config command!\n")

	return nil
}

func CmdEstimate(c *cli.Context) error {
	fmt.Fprint(os.Stdout, "estimate command!\n")

	return nil
}
