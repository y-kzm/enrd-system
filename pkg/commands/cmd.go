package commands

import (
	"github.com/urfave/cli/v2"

	"github.com/y-kzm/enrd-system/cmd/controller/app"
)

var Commands = []*cli.Command{
	commandInit,
	commandTemp,
	commandConf,
	commandEstimate,
}

var commandTemp = &cli.Command{
	Name:   "template",
	Usage:  "Generate tinet config template file",
	Action: app.CmdTemp,
}

var commandInit = &cli.Command{
	Name:   "init",
	Usage:  "Initialize the database",
	Action: app.CmdInit,
}

var commandConf = &cli.Command{
	Name:   "config",
	Usage:  "configure Node from tinet config file",
	Action: app.CmdConf,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "config.yaml",
		},
	},
}

var commandEstimate = &cli.Command{
	Name:   "estimate",
	Usage:  "start to estimate",
	Action: app.CmdEstimate,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "param",
			Aliases: []string{"c"},
			Usage:   "Specify the Parameter file.",
			Value:   "param.yaml",
		},
	},
}
