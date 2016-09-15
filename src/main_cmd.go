package main

import (
	"fmt"
	"fuzzer"

	"github.com/urfave/cli"
)

func (main *MainInterface) Help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func (main *MainInterface) Generate(c *cli.Context) {
	fmt.Println("generate")
}

func (main *MainInterface) Exit(c *cli.Context) {
	main.running = false
}

func (main *MainInterface) Start() {
	main.running = true
}

func (main *MainInterface) Stop() {
	main.running = false
}

func (main *MainInterface) IsRunning() bool {
	return main.running
}

func (main *MainInterface) Encode(c *cli.Context) {
	str := c.Args().Get(0)
	sEnc := fuzzer.Encode(str)
	fmt.Println(sEnc)
}

func (main *MainInterface) Decode(c *cli.Context) {
	str := c.Args().Get(0)
	sDec := fuzzer.Decode(str)
	fmt.Println(sDec)
}

func (main *MainInterface) StartBash(c *cli.Context) {
	Global.BashSession = true
	main.Stop()
}
