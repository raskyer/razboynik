package app

import "github.com/urfave/cli"

func getCommands(app *AppInterface) []cli.Command {
	var helpDefinition = cli.Command{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Help of application",
		Action:  app.Help,
	}

	var runDefinition = cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run reverse shell with configuration",
		Action:  app.Start,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
			cli.StringFlag{Name: "m, method", Usage: "Method to use. Ex: -m POST", Value: "GET"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "fuzzer"},
			cli.IntFlag{Name: "sh, shellmethod", Usage: "Shellmethod to use.", Value: 0},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection", Value: "FromRussiaWithLove<3"},
			cli.BoolFlag{Name: "c, crypt", Usage: "Use a crypt"},
		},
	}

	var generateDefinition = cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate php file",
		Action:  app.Generate,
	}

	return []cli.Command{
		runDefinition,
		generateDefinition,
		helpDefinition,
	}
}
