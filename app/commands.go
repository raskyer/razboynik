package app

import "github.com/urfave/cli"

func getCommands(app *AppInterface) []cli.Command {
	var runDefinition = cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run reverse shell with configuration",
		Action:  app.Start,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
			cli.StringFlag{Name: "m, method", Usage: "Method to use. Ex: -m POST", Value: "GET"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.IntFlag{Name: "s, shellmethod", Usage: "Shellmethod to use. Ex: -s 0", Value: 0},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection", Value: "FromRussiaWithLove<3"},
			cli.BoolFlag{Name: "r, raw", Usage: "Tell razboy if the request should be send raw (without base64 encoding)"},
			cli.BoolFlag{Name: "c, crypt", Usage: "(Not available) Use a crypt"},
		},
	}

	var generateDefinition = cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate php file",
		Action:  app.Generate,
	}

	var scanDefinition = cli.Command{
		Name:    "scan",
		Aliases: []string{"s"},
		Usage:   "Scan a website",
		Action:  app.Scan,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection", Value: "FromRussiaWithLove<3"},
		},
	}

	var invisibleDefinition = cli.Command{
		Name:    "invisible",
		Aliases: []string{"i"},
		Usage:   "Invisible usage",
		Action:  app.Invisible,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the server. Ex: -u http://localhost"},
			cli.StringFlag{Name: "r, referer", Usage: "Url the server will call"},
		},
	}

	var helpDefinition = cli.Command{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Help of application",
		Action:  app.Help,
	}

	return []cli.Command{
		runDefinition,
		generateDefinition,
		scanDefinition,
		invisibleDefinition,
		helpDefinition,
	}
}
