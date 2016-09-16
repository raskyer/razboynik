package app

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/reader"
	"strings"

	"github.com/urfave/cli"
)

func (main *MainInterface) SendRawPHP(c *cli.Context) {
	cmd := strings.Join(c.Args(), " ")
	fuzzer.NET.Send(cmd, reader.Read)
}

func (main *MainInterface) SendLs(c *cli.Context) {
	var ls string

	if c.Bool("raw") {
		r := "ls " + strings.Join(c.Args(), " ")
		ls = fuzzer.CMD.Raw(r)
	} else {
		context := c.Args().Get(0)
		ls = fuzzer.CMD.Ls(context)
	}

	fuzzer.NET.Send(ls, reader.ReadEncode)
}

func (main *MainInterface) SendTest(c *cli.Context) {
	t := "echo 1;"
	fuzzer.NET.Send(t, main.ReceiveTest)
}

func (main *MainInterface) ReceiveTest(str string) {
	if str == "1" {
		fmt.Println("Connected")
		return
	}

	fmt.Println("Error")
}

func (main *MainInterface) ServerSetup(c *cli.Context) {
	srv := &fuzzer.NET

	if c.Bool("i") {
		fmt.Println("url : " + srv.GetUrl())
		fmt.Println("method : " + srv.GetMethodStr())
		fmt.Println("parameter : " + srv.GetParameter())
		return
	}

	url := c.String("u")
	method := c.Int("m")
	parameter := c.String("p")
	crypt := false

	if url == "" && srv.GetUrl() == "" {
		fmt.Println("Flag -u (url) is required")
		return
	}

	if method > 4 {
		fmt.Println("Method is between 0 (default) and 4. [0 => GET, 1 => POST, 3 => HEADER, 4 => COOKIE]")
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
