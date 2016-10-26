package app

import (
	"errors"
	"time"

	"github.com/eatbytes/fuzz/core"
	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/php"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/modules"
	"github.com/eatbytes/fuzzer/printer"
	"github.com/urfave/cli"
)

type MainInterface struct {
	cmd    *cli.App
	server *network.NETWORK
}

func CreateApp() *MainInterface {
	var main MainInterface
	var app *cli.App

	main = MainInterface{}
	app = createCli(&main)
	main.cmd = app

	printer.Intro()

	return &main
}

func createCli(main *MainInterface) *cli.App {
	var client *cli.App

	client = cli.NewApp()

	client.Commands = getCommands(main)
	client.Name = "Fuzzer"
	client.Usage = "File upload meterpreter"
	client.Version = "4.0.0"
	client.Compiled = time.Now()
	client.Authors = []cli.Author{
		cli.Author{
			Name:  "Kamikaze",
			Email: "leakleass@protomail.com",
		},
	}
	client.Copyright = "(c) 2016 Eat Bytes"
	client.EnableBashCompletion = true
	client.BashComplete = func(c *cli.Context) {
		//To edit
		//dm.Fprintf(c.App.Writer, "lipstick\nkiss\nme\nlipstick\nringo\n")
	}
	client.Action = main.Default

	return client
}

func (main *MainInterface) Run(command []string) {
	main.cmd.Run(command)
}

func (main *MainInterface) Help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func (main *MainInterface) Generate(c *cli.Context) {
	printer.Generating()
}

func (main *MainInterface) Start(c *cli.Context) {
	var cf *core.Config
	var err error

	cf, err = main.GetConfig(c)

	if err != nil {
		printer.Error(err)
		return
	}

	err = main.startProcess(cf)

	if err != nil {
		printer.Error(err)
	}
}

func (main *MainInterface) startProcess(cf *core.Config) error {
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

	printer.PrintSection("Reverse shell", "Reverse shell ready!")

	bsh = bash.CreateApp(srv, shl, ph)
	modules.Boot(bsh)
	bsh.Start()

	return nil
}

func (main *MainInterface) GetConfig(c *cli.Context) (*core.Config, error) {
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
