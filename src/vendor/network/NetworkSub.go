package network

import (
	"fmt"
	"worker"

	"github.com/urfave/cli"
)

func requestInfo(c *cli.Context) {
	flag := false
	r := NET._lastResponse.Request
	r.PostForm = NET._body

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
	r := NET._lastResponse
	flag := false

	if c.Bool("status") {
		fmt.Println(r.Status)
		flag = true
	}

	if c.Bool("body") {
		body := worker.GetBody(r)
		fmt.Println("body: %v", string(body))

		flag = true
	}

	if c.Bool("headers") {
		fmt.Println(r.Header)
		flag = true
	}

	if c.Bool("request") {
		fmt.Println(r.Request)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}
}

func Info(c *cli.Context) {
	if NET._lastResponse == nil {
		fmt.Println("You havn't made a request. You must make a request before seeing any information")
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

	fmt.Println("This item : " + item + " is not specified please refer to srv info -h for more info")
}
