package app

import "github.com/urfave/cli"

func getCommands(main *MainInterface) []cli.Command {
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
		Name:    "start",
		Aliases: []string{"s"},
		Usage:   "Start bash",
		Action:  main.Start,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
			cli.IntFlag{Name: "m, method", Usage: "Method to use. Ex: -m 1"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test"},
			cli.BoolFlag{Name: "f, file", Usage: "Use a config from file (default path : ./config)"},
			cli.BoolFlag{Name: "c, crypt", Usage: "Use a crypt"},
			cli.BoolFlag{Name: "i, info", Usage: "Give info on actual config"},
		},
	}

	return []cli.Command{
		generateDefinition,
		startDefinition,
		helpDefinition,
	}
}
