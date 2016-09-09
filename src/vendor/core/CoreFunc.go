package core

import (
	"fmt"

	"github.com/urfave/cli"
)

func help(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func generate(c *cli.Context) {
	fmt.Println("generate")
}

func exit(c *cli.Context) {
	Running = false
}
