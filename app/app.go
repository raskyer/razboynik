package app

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/leaklessgfy/fuzzer/bash"
	"github.com/leaklessgfy/fuzzer/networking"

	"github.com/eatbytes/fuzzcore"
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
	color.Yellow("Generating...")
}

func (main *MainInterface) Start(c *cli.Context) {
	color.Green("Starting...")
	color.Yellow("Trying to communicate with server\n")

	connect := main.ServerSetup(c)
	if !connect {
		return
	}

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
		color.Red("Flag -u (url) is required")

		return false
	}

	if method > 3 {
		color.Red("Method is between 0 (default) and 3.")
		color.Red("[0 => GET, 1 => POST, 2 => HEADER, 3 => COOKIE]")

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
		err.Error()
		return false
	}

	result = fuzzcore.Decode(result)

	if result != "1" {
		color.Red("An error occured with the host")
		fmt.Println(result)

		return false
	}

	color.Green("Successfull connexion")
	return true
}
