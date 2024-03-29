package app

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/y-kzm/enrd-system/api"
	"github.com/y-kzm/enrd-system/pkg/shell"
)

const database = "enrd:PASSWORD@tcp(controller:3306)/enrd"

const port = 52000
const pathTable = "path_info"

type PathInfo struct {
	id   string
	path string
}

type TempInfo struct {
	key   string
	value float64
}

type SelectInfo struct {
	estimation float64
	date       time.Time
}

type ResInfo struct {
	key        string
	value      float64
	first_date time.Time
	last_date  time.Time
}

// Send configuration infomation
func ConfigureRequest(host string, sr []*api.SRInfo) error {
	conn, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
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
		return err
	}

	fmt.Printf("Received from server: Status %d Msg %s\n", r.GetStatus(), r.GetMsg())
	if r.GetStatus() != 0 {
		return fmt.Errorf("%s", r.GetMsg())
	}

	return nil
}

// Send measurement request
func MeasureRequest(host string, method string, param *api.Param) error {
	conn, err := grpc.Dial(host+":"+strconv.Itoa(port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	c := api.NewServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
	defer cancel()

	r, err := c.Measure(ctx, &api.MeasureRequest{
		Method: method,
		Param:  param,
	})
	if err != nil {
		return err
	}

	fmt.Printf("Received from server: Status %d Msg %s\n", r.GetStatus(), r.GetMsg())
	if r.GetStatus() != 0 {
		return fmt.Errorf("%s", r.GetMsg())
	}

	return nil
}

// Parsing yaml files
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

// Temp command
func CmdTemp(c *cli.Context) error {
	// TODO: パス依存
	if err := PrintTemplate("config.yaml"); err != nil {
		return err
	}
	if err := PrintTemplate("param.yaml"); err != nil {
		return err
	}

	return nil
}

// Init command
func CmdInit(c *cli.Context) error {
	fmt.Printf("***** Init Command *****\n")
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	s.Start()
	// time.Sleep(4 * time.Second)

	db, err := sql.Open("mysql", database)
	if err != nil {
		return err
	}
	defer db.Close()

	// Delete all tables
	res, err := db.Query("SHOW TABLES")
	if err != nil {
		return err
	}
	var table string
	for res.Next() {
		res.Scan(&table)
		_, err = db.Exec("DROP TABLE IF EXISTS " + table)
		if err != nil {
			return err
		}
	}
	defer res.Close()

	// Creation of path_info table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + pathTable + " ( id varchar(40) PRIMARY KEY, path varchar(64) ) ")
	if err != nil {
		return err
	}
	fmt.Printf("Done\n")
	s.Stop()

	return nil
}

// Conf command
func CmdConf(c *cli.Context) error {
	fmt.Printf("***** Config Command *****\n")
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	go s.Start()
	// time.Sleep(4 * time.Second)

	// Parsing yaml files
	erconfig, _, err := LoadCfgStruct(c, "config")
	if err != nil {
		return err
	}
	// fmt.Println(erconfig.Config.Rules)  // debug

	// Memorize mapping between host name and SID
	pair := make(map[string]string)
	for _, i := range erconfig.Nodes {
		// pair[i.Host] = i.SID
		pair[i.Host] = i.Locator
	}

	// Check for existing table
	db, err := sql.Open("mysql", database)
	if err != nil {
		return err
	}
	defer db.Close()
	res, _ := db.Query("SELECT * FROM " + pathTable)
	for res.Next() {
		if res.Scan() != nil {
			return err
		}
	}
	defer res.Close()

	path := []*PathInfo{}
	sr := []*api.SRInfo{}

	// Register records for measurement paths
	for i := range erconfig.Config.Rules {
		uuid, err := uuid.NewRandom()
		if err != nil {
			return err
		}
		uuid_str := uuid.String()

		// Create an array with additional source and destination nodes
		path_arr := erconfig.Config.Rules[i].TransitNodes                     // [ Compute2, Compute3 ]
		path_arr = append([]string{erconfig.Config.SrcNode}, path_arr[0:]...) // [ Compute1, Compute2, Compute3 ]
		path_arr = append(path_arr, erconfig.Config.Rules[i].DstNode)         // [ Compute1, Compute2, Compute3, Compute4 ]
		path_str := strings.Join(path_arr, "_")                               // "Compute1_Compute2_Compute3_Compute4"

		path = append(path, &PathInfo{uuid_str, path_str})

		// Insert Execution
		ins, err := db.Prepare("INSERT INTO " + pathTable + " ( id, path ) VALUES( ?, ? )")
		if err != nil {
			return err
		}
		defer ins.Close()
		_, err = ins.Exec(path[i].id, path[i].path)
		if err != nil {
			return err
		}

		// Create a measurement result table for each measurement path
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + path_str + " ( num_meas int unsigned PRIMARY KEY, estimation float, time_stamp datetime ) ")
		if err != nil {
			return err
		}

		// Convert host name to SID (IPv6 address format)
		sid_list := []string{}
		for j := range erconfig.Config.Rules[i].TransitNodes {
			sid_list = append(sid_list, pair[erconfig.Config.Rules[i].TransitNodes[j]])
		}

		sr = append(sr, &api.SRInfo{
			SrcAddr:   erconfig.Config.Rules[i].SrcAddr,
			Vrf:       erconfig.Config.Rules[i].VRF,
			DstAddr:   pair[erconfig.Config.Rules[i].DstNode],
			SidList:   sid_list,
			TableName: path_str,
		})
	}

	if err := ConfigureRequest(erconfig.Config.SrcNode, sr); err != nil {
		return err
	}
	s.Stop()

	return nil
}

// Estimate command
func CmdEstimate(c *cli.Context) error {
	fmt.Printf("***** Estimate command *****\n")
	s := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
	go s.Start()
	// time.Sleep(4 * time.Second)

	// Parse yaml and store in structure array
	_, param, err := LoadCfgStruct(c, "param")
	if err != nil {
		return err
	}
	// fmt.Println(param.Param) // debug

	pm := api.Param{
		PacketNum:   param.Param.PacketNum,
		PacketSize:  param.Param.PacketSize,
		RepeatNum:   param.Param.RepeatNum,
		MeasNum:     param.Param.MeasNum,
		SmaInterval: param.Param.SmaInterval,
	}

	if err := MeasureRequest(param.Param.SrcNode, param.Param.Method, &pm); err != nil {
		return err
	}

	// TODO: 指定の区間の移動平均を計算して標準出力
	s.Stop()

	// Get path information
	db, err := sql.Open("mysql", database+"?parseTime=true")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer db.Close()

	rows, err := db.Query("SELECT path FROM " + pathTable)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	var path_info []string
	for rows.Next() {
		var tmp string
		err := rows.Scan(&tmp)
		if err != nil {
			fmt.Println(err)
			return err
		}
		path_info = append(path_info, tmp)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Loop for path information
	res := []ResInfo{}
	for _, j := range path_info { // [ compute1_compute2_compute4, compute1_compute3_compute4, ... ]
		// Get number of data
		var num_record int
		rows, err := db.Query("SELECT COUNT(*) FROM " + j)
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer rows.Close()
		for rows.Next() {
			err := rows.Scan(&num_record)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
		err = rows.Err()
		if err != nil {
			fmt.Println(err)
			return err
		}
		// Update moving average interval
		if num_record < int(pm.SmaInterval) {
			pm.SmaInterval = int32(num_record)
		}

		// Get data for moving average interval
		slct := []SelectInfo{}
		rows, err = db.Query("SELECT estimation, time_stamp FROM " + j + " ORDER BY time_stamp DESC LIMIT " + strconv.Itoa(int(pm.SmaInterval)))
		if err != nil {
			fmt.Println(err)
			return err
		}
		defer rows.Close()
		for rows.Next() {
			tmp := SelectInfo{}
			err := rows.Scan(&tmp.estimation, &tmp.date)
			if err != nil {
				fmt.Println(err)
				return err
			}
			slct = append(slct, tmp)
		}
		err = rows.Err()
		if err != nil {
			fmt.Println(err)
			return err
		}

		// Calculate moving averages
		var sum float64
		var first_date, last_date time.Time
		first_date = slct[0].date
		last_date = slct[0].date
		sum = 0
		for _, v := range slct {
			sum += v.estimation
			if v.date.Before(first_date) {
				first_date = v.date
			}
			if v.date.After(last_date) {
				last_date = v.date
			}
		}
		res = append(res, ResInfo{
			key:        j,
			value:      sum / float64(len(slct)),
			first_date: first_date,
			last_date:  last_date,
		})
	}

	// Finally, stdout the results
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Path", "Estimation", "First_Date", "Last_Date"})
	max := TempInfo{
		key:   "",
		value: 0,
	}
	for _, j := range res {
		if max.value < j.value {
			max.value = j.value
			max.key = j.key
		}
		var v []string
		v = append(v, j.key)
		v = append(v, strconv.FormatFloat(j.value, 'f', 2, 64))
		v = append(v, j.first_date.Format("2006-01-02 03:04:05"))
		v = append(v, j.last_date.Format("2006-01-02 03:04:05"))
		table.Append(v)
	}
	table.Render()

	fmt.Printf("Note:    Estimation are averages of the %d recent data \n", pm.SmaInterval)
	fmt.Printf("Suggest: %s %.2f \n", max.key, max.value)

	return nil
}

// Display of template files
func PrintTemplate(filename string) error {
	fmt.Println("------------------- " + filename + "\n")
	f, err := os.Open("./templates/" + filename)
	if err != nil {
		return err
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
