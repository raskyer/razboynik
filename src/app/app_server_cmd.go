package app

import (
	"fmt"
	"fuzzer"

	"github.com/urfave/cli"
)

func (main *MainInterface) ServerInfo(c *cli.Context) {
	if fuzzer.NET.GetResponse() == nil {
		fmt.Println("You havn't made any request. You must make a request before being able to see information")
		return
	}

	if len(c.Args()) < 1 {
		fmt.Println("Request => ")
		main.RequestInfo(c)

		fmt.Println("--- >< ---")

		fmt.Println("Response => ")
		main.ResponseInfo(c)

		return
	}
}

func (main *MainInterface) RequestInfo(c *cli.Context) {
	if fuzzer.NET.GetResponse() == nil {
		fmt.Println("You havn't made any request. You must make a request before being able to see information")
		return
	}

	flag := false
	r := fuzzer.NET.GetRequest()

	if c.Bool("url") {
		fmt.Println(r.URL)
		flag = true
	}

	if c.Bool("method") {
		fmt.Println(r.Method)
		flag = true
	}

	if c.Bool("body") {
		fmt.Println(r.PostForm)
		flag = true
	}

	if c.Bool("headers") {
		fmt.Println(r.Header)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}
}

func (main *MainInterface) ResponseInfo(c *cli.Context) {
	if fuzzer.NET.GetResponse() == nil {
		fmt.Println("You havn't made any request. You must make a request before being able to see information")
		return
	}

	flag := false
	r := fuzzer.NET.GetResponse()

	if c.Bool("status") {
		fmt.Println(r.Status)
		flag = true
	}

	if c.Bool("body") {
		body := fuzzer.NET.GetBodyStr(r)
		fmt.Println("body: " + body)

		flag = true
	}

	if c.Bool("headers") {
		fmt.Println(r.Header)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}
}
