package core

import (
	"command"
	"fmt"
	"network"

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
	RunningBash = false
	RunningMain = true
}

func exit(c *cli.Context) {
	Running = false
}
