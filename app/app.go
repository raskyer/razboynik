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
		fmt.Fprintf(c.App.Writer, "run\nscan\ngenerate\n")
	}

	client.Flags = []cli.Flag{
		cli.BoolFlag{Name: "help, h", Usage: "Help"},
	}

	app.cmd = client
}

func (app *AppInterface) Run(command []string) {
	app.cmd.Run(command)
}

func (app *AppInterface) Help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func (app *AppInterface) Generate(c *cli.Context) {
	services.PrintGenerating()
}
