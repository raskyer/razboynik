package app

import (
	"time"

	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/bash/networking"
	"github.com/eatbytes/fuzzer/printer"
	"github.com/urfave/cli"
)

type MainInterface struct {
	cmd *cli.App
}

func CreateMainApp() *MainInterface {
	main := MainInterface{}
	app := createCli(&main)
	main.cmd = app

	return &main
}

func createCli(main *MainInterface) *cli.App {
	client := cli.NewApp()
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
	printer.Start()

	connect := main.ServerSetup(c)
	if !connect {
		return
	}

	printer.End()

	bsh := bash.CreateBashApp()
	bsh.Start()
}

func (main *MainInterface) ServerSetup(c *cli.Context) bool {
	srv := &fuzzcore.NET

	url := c.String("u")
	method := c.Int("m")
	parameter := c.String("p")
	crypt := false

	if url == "" {
		printer.SetupError(0)
		return false
	}

	if method > 3 {
		printer.SetupError(1)
		return false
	}

	if parameter == "" {
		parameter = srv.GetParameter()
	}

	srv.SetConfig(url, method, parameter, crypt)

	return main.SendTest(c)
}

func (main *MainInterface) SendTest(c *cli.Context) bool {
	t := fuzzcore.PHP.Raw("$r=1;")
	result, err := networking.Process(t)

	if err != nil {
		printer.Error(err)
		return false
	}

	result = fuzzcore.Decode(result)

	if result != "1" {
		printer.Test(false, result)
		return false
	}

	printer.Test(true, result)
	return true
}
