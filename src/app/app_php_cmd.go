package app

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/common"
	"strings"

	"github.com/urfave/cli"
)

func (main *MainInterface) SendRawPHP(c *cli.Context) {
	cmd := strings.Join(c.Args(), " ")
	raw := fuzzer.PHP.Raw(cmd)
	r, err := common.Send(raw)

	if err != nil {
		err.Error()
		return
	}

	var resultStr string

	m := c.Int("r")

	if m != fuzzer.NET.GetMethod() && m < 4 {
		resultStr = fuzzer.NET.GetResultStrByMethod(m, r)
	} else {
		resultStr = fuzzer.NET.GetResultStr(r)
	}

	if c.Bool("d") {
		//Download logic
	}

	common.Read(resultStr)
}

func (main *MainInterface) SendTest(c *cli.Context) {
	t := fuzzer.PHP.Raw("$r=1;")
	result, err := common.Process(t)

	if err != nil {
		err.Error()
		return
	}

	result = fuzzer.Decode(result)

	if result != "1" {
		fmt.Println("An error occured with the host")
		fmt.Println(result)

		prompt := "\033[31m•\033[0m\033[31m»\033[0m "
		main.SetPrompt(prompt)
		return
	}

	fmt.Println("Connected")
	prompt := "\033[32m•\033[0m\033[32m»\033[0m "
	main.SetPrompt(prompt)
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
		fmt.Println("Please write the path of the local file to download")
		return
	}

	path := arr[0]
	loc := "output.txt"

	if len(arr) > 1 {
		loc = arr[1]
	}

	common.Download(path, loc)
}
