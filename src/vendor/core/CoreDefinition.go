package core

import (
	"command"
	"network"

	"github.com/urfave/cli"
)

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

var srvDefinition = cli.Command{
	Name:    "srv",
	Aliases: []string{"s"},
	Usage:   "Prefix to make server command",
	Subcommands: []cli.Command{
		{
			Name:   "bash",
			Usage:  "Open meterpreter like session (command are send raw to server except 'cd')",
			Action: bash,
		},
		{
			Name:   "ls",
			Usage:  "List file on server",
			Action: command.CMD.Ls,
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
				cli.IntFlag{Name: "m, method", Usage: "Method to use. Ex: -m 1"},
				cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test"},
				cli.BoolFlag{Name: "f, file", Usage: "Use a config from file (default path : ./config)"},
				cli.BoolFlag{Name: "c, crypt", Usage: "Use a crypt"},
			},
		},
		{
			Name:   "info",
			Usage:  "Give information on the last specified item. Ex: srv info request",
			Action: network.Info,
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "url", Usage: "Shows request's url"},
				cli.BoolFlag{Name: "method", Usage: "Shows request's method"},
				cli.BoolFlag{Name: "status", Usage: "Shows response's status"},
				cli.BoolFlag{Name: "request", Usage: "Shows response's interpreted request"},
				cli.BoolFlag{Name: "body", Usage: "Shows item's body"},
				cli.BoolFlag{Name: "headers", Usage: "Shows item's headers"},
			},
		},
	},
}
