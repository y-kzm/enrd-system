package app

// TODO: log.Faitalの廃止

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/y-kzm/enrd-system/api"
	"github.com/y-kzm/enrd-system/pkg/shell"
)

// Changed in production environment
const database = "enrd:0ta29SourC3@tcp(controller:3306)/enrd"
//const database = "enrd:0ta29SourC3@tcp(localhost:3306)/enrd"
const port = 52000
const pathTable = "path_info"

type PathInfo struct {
	id   string
	path string
	num  int
}

// pb.goで定義されてるっぽい？
type SRInfo struct {
	src_addr   string
	vrf        int32
	dst_addr   string
	sid_list   []string
	table_name string
}


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
	// TODO: パスに依存しない形にする
	PrintTemplate("config.yaml")
	PrintTemplate("param.yaml")

	return nil
}

func CmdInit(c *cli.Context) error {
	fmt.Fprint(os.Stdout, "***** Init Command! *****\n")

	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	/* Delete all tables */
	res, _ := db.Query("SHOW TABLES")
	var table string
	for res.Next() {
		res.Scan(&table)
		_, err = db.Exec("DROP TABLE IF EXISTS " + table)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer res.Close()

	/* Creation of path_info table */
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + pathTable + " ( id varchar(40) PRIMARY KEY, path varchar(64), num smallint unsigned ) ")
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func CmdConf(c *cli.Context) error {
	fmt.Fprint(os.Stdout, "***** Config Command! *****\n")

	/* Parsing yaml files */
	erconfig, _, err := LoadCfgStruct(c, "config")
	if err != nil {
		return err
	}
	// fmt.Println(erconfig.Config.Rules)

	/* Memorize mapping between host name and SID */
	pair := make(map[string]string)
	for _, i := range erconfig.Nodes {
		pair[i.Host] = i.SID
	}

	/* Check for existing table */
	db, err := sql.Open("mysql", database)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	res, _ := db.Query("SELECT * FROM " + pathTable)
	for res.Next() {
		if res.Scan() != nil {
			log.Fatal("database is already exist: ", err)
		}
	}
	defer res.Close()

	path := []*PathInfo{}
	sr := []*api.SRInfo{}

	// Register records for measurement paths
	for i := range erconfig.Config.Rules {
		uuid, err := uuid.NewRandom()
		if err != nil {
			log.Fatal(err)
		}
		uuid_str := uuid.String()

		/* Create an array with additional source and destination nodes */
		path_arr := erconfig.Config.Rules[i].TransitNodes                     // [ Compute2, Compute3 ]
		path_arr = append([]string{erconfig.Config.SrcNode}, path_arr[0:]...) // [ Compute1, Compute2, Compute3 ]
		path_arr = append(path_arr, erconfig.Config.Rules[i].DstNode)         // [ Compute1, Compute2, Compute3, Compute4 ]
		path_str := strings.Join(path_arr, "_")                               // "Compute1->Compute2->Compute3->Compute4"

		path = append(path, &PathInfo{uuid_str, path_str, 0})

		/* Insert Execution */
		ins, err := db.Prepare("INSERT INTO " + pathTable + " (id,path,num) VALUES(?,?,?)")
		if err != nil {
			log.Fatal(err)
		}
		defer ins.Close()
		_, err = ins.Exec(path[i].id, path[i].path, path[i].num)
		if err != nil {
			log.Fatal(err)
		}

		/* Create a measurement result table for each measurement path */
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + path_str + " ( cycle int unsigned PRIMARY KEY, estimate float, timestamp datetime ) ")
		if err != nil {
			log.Fatal(err)
		}

		/* Convert host name to SID (IPv6 address format) */
		sid_list := []string{}
		for j := range erconfig.Config.Rules[i].TransitNodes {
			sid_list = append(sid_list, pair[erconfig.Config.Rules[i].TransitNodes[j]])
		}


		sr = append(sr, &api.SRInfo{
			SrcAddr: erconfig.Config.Rules[i].SrcAddr,
			Vrf: erconfig.Config.Rules[i].VRF,
			DstAddr: pair[erconfig.Config.Rules[i].DstNode],
			SidList: sid_list, 
			TableName: path_str,
		})
	}

	// TODO: SRInfoを構成する
	// ConfigureReq()の引数で渡してあげる
	//var api.SRInfo []sr_info

	// それをそのままgrpcで送る
	ConfigureRequest(erconfig.Config.SrcNode, sr)

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

	// TODO: パス情報テーブルからレコードを検索
	// 存在しない場合はエラー
	// 存在すればOK & 計測テーブルここで作る？

	// gRPCフォーマットにテーブル名(UUID)を追加する必要がある  ****重要****
	// path_infoから検索してくる必要あり "->"でいい感じに連結してくれる関数を作るべき
	// https://stackoverflow.com/questions/36344826/how-do-i-represent-a-uuid-in-a-protobuf-message

	return nil
}

func PrintTemplate(filename string) {
	fmt.Print("------------------- " + filename + "\n")
	f, err := os.Open("./templates/" + filename)
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

func ConfigureRequest(addr string, sr []*api.SRInfo) {
	conn, err := grpc.Dial(addr+":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := api.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Configure(ctx, &api.ConfigureRequest{
		Msg:    "go",
		SrInfo: sr,
	})
	if err != nil {
		log.Fatalf("Could not echo: %v", err)
	}

	// TODO: 戻り値チェック
	fmt.Printf("Received from server: Status %d Msg %s\n", r.GetStatus(), r.GetMsg())
}


