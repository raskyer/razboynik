package app

import "github.com/urfave/cli"

func _getCommands(main *MainInterface) []cli.Command {
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

	var startDefinition = cli.Command{
		Name:   "start",
		Usage:  "Start bash",
		Action: main.Start,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
			cli.IntFlag{Name: "m, method", Usage: "Method to use. Ex: -m 1"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test"},
			cli.BoolFlag{Name: "f, file", Usage: "Use a config from file (default path : ./config)"},
			cli.BoolFlag{Name: "c, crypt", Usage: "Use a crypt"},
			cli.BoolFlag{Name: "i, info", Usage: "Give info on actual config"},
		},
	}

	var srvDefinition = cli.Command{
		Name:    "srv",
		Aliases: []string{"s"},
		Usage:   "Prefix to make server command",
		Subcommands: []cli.Command{
			{
				Name:   "info",
				Usage:  "Give information on the last specified item. Ex: srv info request",
				Action: main.ServerInfo,
				Subcommands: []cli.Command{
					{
						Name:   "response",
						Action: main.ResponseInfo,
						Flags: []cli.Flag{
							cli.BoolFlag{Name: "status", Usage: "Shows response's status"},
							cli.BoolFlag{Name: "body", Usage: "Shows item's body"},
							cli.BoolFlag{Name: "headers", Usage: "Shows item's headers"},
						},
					},
					{
						Name:   "request",
						Action: main.RequestInfo,
						Flags: []cli.Flag{
							cli.BoolFlag{Name: "url", Usage: "Shows request's url"},
							cli.BoolFlag{Name: "method", Usage: "Shows request's method"},
							cli.BoolFlag{Name: "body", Usage: "Shows item's body"},
							cli.BoolFlag{Name: "headers", Usage: "Shows item's headers"},
						},
					},
				},
				Flags: []cli.Flag{
					cli.BoolFlag{Name: "url", Usage: "Shows request's url"},
					cli.BoolFlag{Name: "method", Usage: "Shows request's method"},
					cli.BoolFlag{Name: "status", Usage: "Shows response's status"},
					cli.BoolFlag{Name: "body", Usage: "Shows item's body"},
					cli.BoolFlag{Name: "headers", Usage: "Shows item's headers"},
				},
			},
		},
	}

	return []cli.Command{
		generateDefinition,
		startDefinition,
		srvDefinition,
		helpDefinition,
	}
}
