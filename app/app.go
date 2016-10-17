package app

import (
	"time"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/normalizer"
	"github.com/eatbytes/fuzz/php"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/modules"
	"github.com/eatbytes/fuzzer/printer"
	"github.com/eatbytes/fuzzer/processor"
	"github.com/urfave/cli"
)

type MainInterface struct {
	cmd    *cli.App
	server *network.NETWORK
}

func CreateMainApp() *MainInterface {
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
	var err error
	var srv *network.NETWORK
	var shl *shell.SHELL
	var bsh *bash.BashInterface

	printer.Start()

	srv, err = main.ServerSetup(c)

	if err != nil {
		printer.Error(err)
		return
	}

	printer.End()

	shl = &shell.SHELL{}
	bsh = bash.CreateBashApp(srv, shl)
	modules.Boot(bsh)

	bsh.Start()
}

func (main *MainInterface) ServerSetup(c *cli.Context) (*network.NETWORK, error) {
	var url, parameter string
	var method int
	var result bool
	var err error
	var srv *network.NETWORK

	url = c.String("u")
	method = c.Int("m")
	parameter = c.String("p")

	if url == "" {
		//printer.SetupError(0)
		return nil, nil
	}

	if method > 3 {
		//printer.SetupError(1)
		return nil, nil
	}

	if parameter == "" {
		parameter = "fuzzer"
	}

	result, err = main.SendTest(c)

	if err != nil {
		return nil, err
	}

	srv.Setup(url, parameter, method)

	return srv, nil
}

func (main *MainInterface) SendTest(c *cli.Context) (bool, error) {
	var r string
	var result string
	var err error

	r = php.Raw("$r=1;")
	result, err = processor.Process(r)

	if err != nil {
		return false, err
	}

	result, err = normalizer.Decode(result)

	if err != nil {
		return false, err
	}

	if result != "1" {
		//printer.Test(false, result)
		return false, nil
	}

	return true, result
}
