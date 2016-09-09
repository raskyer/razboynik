package core

import (
	"fmt"
	"network"

	"command"

	"github.com/urfave/cli"
)

var CMD = command.CMD{}
var NET = network.NET

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
				cli.BoolFlag{Name: "raw"},
			},
		},
		{
			Name:   "config",
			Usage:  "configure server",
			Action: NET.Setup,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "u"},
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

var template = cli.Command{
	Name:    "template",
	Aliases: []string{"t"},
	Usage:   "options for task templates",
	Subcommands: []cli.Command{
		{
			Name:  "add",
			Usage: "add a new template",
			Action: func(c *cli.Context) error {
				fmt.Println("new task template: ", c.Args().First())
				return nil
			},
		},
		{
			Name:  "remove",
			Usage: "remove an existing template",
			Action: func(c *cli.Context) error {
				fmt.Println("removed task template: ", c.Args().First())
				return nil
			},
		},
	},
}
