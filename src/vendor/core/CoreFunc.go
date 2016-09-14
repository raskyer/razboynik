package core

import (
	"command"
	"fmt"
	"network"
	"normalizer"

	"github.com/urfave/cli"
)

func help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func generate(c *cli.Context) {
	fmt.Println("generate")
}

func bash(c *cli.Context) {
	if !network.NET.IsSetup() {
		command.HandleNotConnected()
		return
	}

	RunningMain = false
	RunningBash = true
}

func ExitBash() {
	command.CMD.Reset()
	RunningBash = false
	RunningMain = true
}

func exit(c *cli.Context) {
	Running = false
}

func encode(c *cli.Context) {
	str := c.Args().Get(0)
	sEnc := normalizer.Encode(str)
	fmt.Println(sEnc)
}

func decode(c *cli.Context) {
	str := c.Args().Get(0)
	sDec := normalizer.Decode(str)
	fmt.Println(sDec)
}
