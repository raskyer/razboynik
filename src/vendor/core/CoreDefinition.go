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
	Usage:   "Generate php file",
	Action:  generate,
}

var exitDefinition = cli.Command{
	Name:    "exit",
	Aliases: []string{"e"},
	Usage:   "Exit the application",
	Action:  exit,
}

var cmdDefinition = cli.Command{
	Name:    "cmd",
	Aliases: []string{"c"},
	Usage:   "Prefix to make cmd command",
	Subcommands: []cli.Command{
		{
			Name:   "config",
			Usage:  "Configure cmd",
			Action: CMD.Setup,
		},
	},
}

var srvDefinition = cli.Command{
	Name:    "srv",
	Aliases: []string{"s"},
	Usage:   "Prefix to make server command",
	Subcommands: []cli.Command{
		{
			Name:   "ls",
			Usage:  "List file on server",
			Action: CMD.Ls,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "raw, send a raw ls"},
			},
		},
		{
			Name:   "config",
			Usage:  "Configure server",
			Action: network.NET.Setup,
			Flags: []cli.Flag{
				cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
				cli.IntFlag{Name: "m, method", Usage: "Method to use. Ex: -m 1 [0 = GET, 1 = POST, 2 = HEADER, 3 = COOKIE]"},
				cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test"},
				cli.BoolFlag{Name: "f, file", Usage: "Use a config from file (default path : ./config). Ex: -f"},
				cli.BoolFlag{Name: "c, crypt", Usage: "Use a crypt. Ex: -c"},
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
