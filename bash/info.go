package bash

import (
	"fmt"
	"strings"

	"github.com/eatbytes/fuzzcore"
	"github.com/fatih/color"
)

func (b *BashInterface) Info(cmd *BashCommand) {
	if fuzzcore.NET.GetResponse() == nil {
		color.Red("You havn't made any request.")
		color.Red("You must make a request before seing any information")
		return
	}

	b.RequestInfo(cmd.str)
	b.ResponseInfo(cmd.str)
}

func (b *BashInterface) RequestInfo(str string) {
	color.Yellow("--- Request ---")

	flag := false
	r := fuzzcore.NET.GetRequest()

	if strings.Contains(str, "-url") {
		color.Cyan("Url: ")
		fmt.Println(r.URL.String())
		flag = true
	}

	if strings.Contains(str, "-method") {
		color.Cyan("Method: ")
		fmt.Println(r.Method)
		flag = true
	}

	if strings.Contains(str, "-body") {
		color.Cyan("Body: ")
		fmt.Println(r.PostForm)
		flag = true
	}

	if strings.Contains(str, "-header") {
		color.Cyan("Header: ")
		fmt.Println(r.Header)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}

	color.Unset()
}

func (b *BashInterface) ResponseInfo(str string) {
	color.Yellow("--- Response ---")

	flag := false
	r := fuzzcore.NET.GetResponse()

	if strings.Contains(str, "-status") {
		color.Cyan("Status:")
		fmt.Println(r.Status)
		flag = true
	}

	if strings.Contains(str, "-body") {
		color.Cyan("Body:")
		body := fuzzcore.NET.GetBodyStr(r)
		fmt.Println(body)
		flag = true
	}

	if strings.Contains(str, "-headers") {
		color.Cyan("Headers:")
		fmt.Println(r.Header)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}

	color.Unset()
}
