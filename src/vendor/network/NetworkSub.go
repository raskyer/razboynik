package network

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli"
)

func RequestInfo(c *cli.Context) {
	if NET._lastResponse == nil {
		fmt.Println("You havn't made a request. You must make a request before seeing any information")
		return
	}

	request := NET._lastResponse.Request
	flag := false
	request.PostForm = NET._body

	if c.Bool("url") {
		showRequestUrl(request)
		flag = true
	}

	if c.Bool("method") {
		showRequestMethod(request)
		flag = true
	}

	if c.Bool("body") {
		showRequestBody(request)
		flag = true
	}

	if c.Bool("headers") {
		showRequestHeaders(request)
		flag = true
	}

	if !flag {
		fmt.Println(request)
	}
}

func ResponseInfo(c *cli.Context) {
	if NET._lastResponse == nil {
		fmt.Println("You havn't made a request. You must make a request before seeing any information")
		return
	}

	response := NET._lastResponse
	flag := false

	if c.Bool("status") {
		showResponseStatus(response)
		flag = true
	}

	if c.Bool("body") {
		showResponseBody(response)
		flag = true
	}

	if c.Bool("headers") {
		showResponseHeaders(response)
		flag = true
	}

	if c.Bool("request") {
		showResponseRequest(response)
		flag = true
	}

	if !flag {
		fmt.Println(response)
	}
}

func SpecifiedItem(c *cli.Context) {
	fmt.Println("Please specified the item: response or request")
}

func showRequestUrl(r *http.Request) {
	fmt.Println(r.URL)
}

func showRequestMethod(r *http.Request) {
	fmt.Println(r.Method)
}

func showRequestBody(r *http.Request) {
	fmt.Println(r.PostForm)
}

func showRequestHeaders(r *http.Request) {
	fmt.Println(r.Header)
}

func showResponseStatus(r *http.Response) {
	fmt.Println(r.Status)
}

func showResponseBody(r *http.Response) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("body: %v", string(body))
}

func showResponseHeaders(r *http.Response) {
	fmt.Println(r.Header)
}

func showResponseRequest(r *http.Response) {
	fmt.Println(r.Request)
}
