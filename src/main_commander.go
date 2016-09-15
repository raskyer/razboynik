package main

import (
	"fmt"
	"fuzzer"
	"strings"

	"github.com/urfave/cli"
)

func (main *MainInterface) SendRawPHP(c *cli.Context) {
	cmd := strings.Join(c.Args(), " ")
	fuzzer.NET.Send(cmd, main.ReceiveRawPHP)
}

func (main *MainInterface) ReceiveRawPHP(result string) {
	fmt.Println(result)
}

func (main *MainInterface) SendLs(c *cli.Context) {
	if !fuzzer.NET.IsSetup() {
		fmt.Println("You haven't setup the required information, please refer to srv config")
		return
	}

	var ls string

	if c.Bool("raw") {
		r := "ls " + strings.Join(c.Args(), " ")
		ls = fuzzer.CMD.Raw(r)
	} else {
		context := c.Args().Get(0)
		ls = fuzzer.CMD.Ls(context)
	}

	fuzzer.NET.Send(ls, main.ReceiveLs)
}

func (main *MainInterface) ReceiveLs(result string) {
	fmt.Println(result)
}

func (main *MainInterface) ServerSetup(c *cli.Context) {
	srv := &fuzzer.NET

	url := c.String("u")
	method := c.Int("m")
	parameter := c.String("p")
	crypt := false

	if url == "" {
		fmt.Println("flag -u (url) is required")
		return
	}

	if parameter == "" {
		parameter = srv.GetParameter()
	}

	srv.SetConfig(url, method, parameter, crypt)
}

func (main *MainInterface) ServerInfo(c *cli.Context) {
	if fuzzer.NET.GetResponse() == nil {
		fmt.Println("You havn't made any request. You must make a request before being able to see information")
		return
	}

	if len(c.Args()) < 1 {
		fmt.Println("Request => ")
		requestInfo(c)

		fmt.Println("Response => ")
		responseInfo(c)

		return
	}

	item := c.Args().Get(0)

	if item == "request" {
		requestInfo(c)
		return
	}

	if item == "response" {
		responseInfo(c)
		return
	}

	fmt.Println("The item : " + item + " is not specified. Please refer to srv info -h for more information")
}
