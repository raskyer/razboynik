package main

import (
	"fmt"
	"fuzzer"

	"github.com/urfave/cli"
)

func requestInfo(c *cli.Context) {
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

func responseInfo(c *cli.Context) {
	flag := false
	r := fuzzer.NET.GetResponse()

	if c.Bool("status") {
		fmt.Println(r.Status)
		flag = true
	}

	if c.Bool("body") {
		body := fuzzer.GetBody(r)
		fmt.Println("body: " + string(body))

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
