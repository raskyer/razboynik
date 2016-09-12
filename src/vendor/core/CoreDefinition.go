package core

import (
	"fmt"
	"network"

	"command"

	"github.com/urfave/cli"
)

var CMD = command.CMD{}

var helpDefinition = cli.Command{
	Name:    "help",
	Aliases: []string{"h"},
	Usage:   "Help of application",
	Action:  help,
	Flags: []cli.Flag{
		cli.BoolFlag{Name: "t"},
	},
}

var generateDefinition = cli.Command{
	Name:    "generate",
	Aliases: []string{"g"},
	Usage:   "generate a bungle",
	Action:  generate,
}

var exitDefinition = cli.Command{
	Name:    "exit",
	Aliases: []string{"e"},
	Usage:   "exit the application",
	Action:  exit,
}

var srvDefinition = cli.Command{
	Name:    "srv",
	Aliases: []string{"s"},
	Usage:   "prefix to make server command",
	Subcommands: []cli.Command{
		{
			Name:   "ls",
			Usage:  "list file on server",
			Action: CMD.Ls,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "raw, send a raw ls"},
			},
		},
		{
			Name:   "config",
			Usage:  "configure server",
			Action: network.NET.Setup,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "u, url of the server"},
				cli.IntFlag{Name: "m, method to use"},
				cli.StringFlag{Name: "p, parameter to use"},
			},
		},
		{
			Name:  "t",
			Usage: "test",
			Action: func(c *cli.Context) error {
				fmt.Println("test")
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:  "t2",
					Usage: "not",
					Action: func(c *cli.Context) {
						fmt.Println("t2")
					},
				},
			},
		},
	},
}
