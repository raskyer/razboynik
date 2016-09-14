package php

import (
	"fmt"
	"network"
	"strings"

	"github.com/urfave/cli"
)

func Raw(c *cli.Context) {
	cmd := strings.Join(c.Args(), " ")

	network.NET.Send(cmd, phpEnd)
}

func phpEnd(resp string) {
	fmt.Println(resp)
}
