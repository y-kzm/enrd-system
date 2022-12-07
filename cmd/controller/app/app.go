package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/y-kzm/enrd-system/api"
	"github.com/y-kzm/enrd-system/pkg/shell"
)

const port = 5200

func LoadCfgStruct(c *cli.Context, filename string) (erconfig shell.ErConfig, erparam shell.ErParam, err error) {
	cfgFile := c.String(filename)
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("yaml")

		if err = viper.ReadInConfig(); err != nil {
			return erconfig, erparam, err
		}
		if filename == "config" {
			if err = viper.Unmarshal(&erconfig); err != nil {
				return erconfig, erparam, err
			}
		} else if filename == "param" {
			if err = viper.Unmarshal(&erparam); err != nil {
				return erconfig, erparam, err
			}
		}
	} else {
		err = fmt.Errorf("not set %s file.", filename)
		return erconfig, erparam, err
	}

	return erconfig, erparam, nil
}

func CmdTemp(c *cli.Context) error {
	// TODO: viperでパスに依存しない形にする
	PrintTemplate("config_template.yaml")
	PrintTemplate("param_template.yaml")

	return nil
}

func CmdInit(c *cli.Context) error {
	fmt.Fprint(os.Stdout, "init command!\n")

	return nil
}

func CmdConf(c *cli.Context) error {
	fmt.Fprint(os.Stdout, "config command!\n")

	// yaml解析して構造体配列に格納
	erconfig, _, err := LoadCfgStruct(c, "config")
	if err != nil {
		return err
	}
	fmt.Println(erconfig.Config.Rules)

	// 対応表の作成OK
	pair := make(map[string]string)
	for _, i := range erconfig.Nodes {
		pair[i.Host] = i.SID
	}

	// それをそのままgrpcで送る
	ConfigureRequest(erconfig.Config.SrcNode)

	return nil
}

func CmdEstimate(c *cli.Context) error {
	fmt.Fprint(os.Stdout, "estimate command!\n")

	// yaml解析して構造体配列に格納
	_, param, err := LoadCfgStruct(c, "param")
	if err != nil {
		return err
	}
	fmt.Println(param.Param) // パラメータ取得OKc

	return nil
}

func PrintTemplate(filename string) {
	fmt.Print("------------------- " + filename + "\n")
	f, err := os.Open("./configs/" + filename)
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
}

func ConfigureRequest(addr string) {
	conn, err := grpc.Dial(addr+":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Configure(ctx, &api.ConfigureRequest{
		Msg:    "hello!",
		SrInfo: nil,
	})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}
	log.Printf("Received from server: %d %s", r.GetStatus(), r.GetMsg())
}
