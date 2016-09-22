package app

import (
	"fmt"
	"fuzzer"

	"github.com/urfave/cli"
)

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

	if method > 3 {
		fmt.Println("Method is between 0 (default) and 3. [0 => GET, 1 => POST, 2 => HEADER, 3 => COOKIE]")
		return
	}

	if parameter == "" {
		parameter = srv.GetParameter()
	}

	srv.SetConfig(url, method, parameter, crypt)
	main.SendTest(c)
}

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
