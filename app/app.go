package app

import (
	"errors"
	"fmt"
	"time"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/php"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
	"github.com/eatbytes/razboynik/modules"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

type AppInterface struct {
	cmd    *cli.App
	server *network.NETWORK
}

func CreateApp() *AppInterface {
	var app *AppInterface
	var cli *cli.App

	services.PrintIntro()

	app = &AppInterface{}
	cli = createCli(app)
	app.cmd = cli

	return app
}

func createCli(app *AppInterface) *cli.App {
	var client *cli.App

	client = cli.NewApp()

	client.Commands = getCommands(app)
	client.Name = "RazBOYNiK"
	client.Usage = "Reverse shell via file upload exploit"
	client.Version = "1.4.0"
	client.Compiled = time.Now()
	client.Authors = []cli.Author{
		cli.Author{
			Name:  "Kamikaze",
			Email: "leakleass@protomail.com",
		},
	}
	client.Copyright = "(c) 2016 EatBytes. из России с любовью <3"
	client.EnableBashCompletion = true
	client.BashComplete = func(c *cli.Context) {
		fmt.Fprintf(c.App.Writer, "start\ngenerate\n")
	}
	client.Action = app.Default
	client.Flags = []cli.Flag{
		cli.BoolFlag{Name: "help, h", Usage: "Help"},
	}

	return client
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

func (app *AppInterface) Start(c *cli.Context) {
	var cf *core.Config
	var err error

	cf, err = app.GetConfig(c)

	if err != nil {
		services.PrintError(err)
		return
	}

	services.PrintStart()

	err = app.startProcess(cf)

	if err != nil {
		services.PrintError(err)
		return
	}
}

func (app *AppInterface) startProcess(cf *core.Config) error {
	var srv *network.NETWORK
	var shl *shell.SHELL
	var ph *php.PHP
	var bsh *bash.BashInterface
	var status bool
	var err error

	srv = &network.NETWORK{}
	ph = &php.PHP{}
	shl = &shell.SHELL{}

	err = srv.Setup(cf)

	if err != nil {
		return err
	}

	status, err = srv.Test()

	if err != nil || status != true {
		return err
	}

	services.PrintSection("Reverse shell", "Reverse shell ready!")

	bsh = bash.CreateApp(srv, shl, ph)
	modules.Boot(bsh)
	bsh.Start()

	return nil
}

func (app *AppInterface) GetConfig(c *cli.Context) (*core.Config, error) {
	var url, parameter, method string
	var shmethod int
	var cf core.Config
	var err error

	url = c.String("u")
	method = c.String("m")
	parameter = c.String("p")

	if url == "" {
		err = errors.New("Flag -u (url) is required")
		return nil, err
	}

	if parameter == "" {
		parameter = "fuzzer"
	}

	cf = core.Config{
		url,
		method,
		parameter,
		shmethod,
		false,
	}

	return &cf, nil
}
