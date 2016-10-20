package modules

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/eatbytes/fuzzer/bash"
	"github.com/fatih/color"
)

func Info(bc *bash.BashCommand) {
	if bc.GetServer().GetResponse() == nil {
		color.Red("You havn't made any request.")
		color.Red("You must make a request before seing any information")
		return
	}

	if !strings.Contains(bc.GetStr(), "response") {
		RequestInfo(bc)
	}

	if !strings.Contains(bc.GetStr(), "request") {
		ResponseInfo(bc)
	}
}

func RequestInfo(bc *bash.BashCommand) {
	var flag bool
	var str string
	var r *http.Request

	color.Yellow("--- Request ---")

	flag = false
	r = bc.GetServer().GetRequest()
	str = bc.GetStr()

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

func ResponseInfo(bc *bash.BashCommand) {
	var flag bool
	var str string
	var r *http.Response

	color.Yellow("--- Response ---")

	flag = false
	r = bc.GetServer().GetResponse()
	str = bc.GetStr()

	if strings.Contains(str, "-status") {
		color.Cyan("Status:")
		fmt.Println(r.Status)
		flag = true
	}

	if strings.Contains(str, "-body") {
		color.Cyan("Body:")
		body := bc.GetServer().GetBodyStr(r)
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
