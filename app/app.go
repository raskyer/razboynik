package app

import (
	"errors"
	"time"

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

	return &main
}

func createCli(main *MainInterface) *cli.App {
	var client *cli.App

	client = cli.NewApp()
	client.Commands = getCommands(main)
	client.CommandNotFound = func(c *cli.Context, command string) {
		cli.ShowAppHelp(c)
	}

	client.Name = "Fuzzer"
	client.Version = "4.0.0"
	client.Compiled = time.Now()
	client.Authors = []cli.Author{
		cli.Author{
			Name:  "Kamikaze",
			Email: "leakleass@protomail.com",
		},
	}
	client.Copyright = "(c) 2016 Eat Bytes"
	client.Usage = "File upload meterpreter"

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
	var srv *network.NETWORK
	var shl *shell.SHELL
	var bsh *bash.BashInterface
	var ph *php.PHP
	var err error
	var status bool

	ph = &php.PHP{}
	shl = &shell.SHELL{}

	printer.Start()

	srv, err = main.ServerSetup(c)

	if err != nil {
		printer.Error(err)
		return
	}

	status, err = srv.Test()

	if err != nil || status != true {
		printer.Error(err)
		return
	}

	printer.End()

	bsh = bash.CreateApp(srv, shl, ph)
	modules.Boot(bsh)

	bsh.Start()
}

func (main *MainInterface) ServerSetup(c *cli.Context) (*network.NETWORK, error) {
	var url, parameter string
	var method int
	var err error
	var srv *network.NETWORK

	srv = &network.NETWORK{}
	url = c.String("u")
	method = c.Int("m")
	parameter = c.String("p")

	if url == "" {
		err = errors.New("Flag -u (url) is required")
		return nil, err
	}

	if method > 3 {
		err = errors.New("Method is between 0 (default) and 3.")
		return nil, err
	}

	if parameter == "" {
		parameter = "fuzzer"
	}

	srv.Setup(url, parameter, method)

	return srv, nil
}
