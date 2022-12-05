package commands

import "github.com/urfave/cli/v2"

var Commands = []*cli.Command{
	commandInit,
	commandConf,
	commandEstimate,
}

var commandInit = &cli.Command{
	Name:   "init",
	Usage:  "Generate tinet config template file",
	Action: CmdInit,
}

var commandConf = &cli.Command{
	Name:   "conf",
	Usage:  "configure Node from tinet config file",
	Action: CmdConf,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config",
			Aliases: []string{"c"},
			Usage:   "Specify the Config file.",
			Value:   "conf.yaml",
		},
	},
}

var commandEstimate = &cli.Command{
	Name:   "estimate",
	Usage:  "start to estimate",
	Action: CmdEstimate,
}
