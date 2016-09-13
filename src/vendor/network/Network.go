package network

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"

	"github.com/urfave/cli"
)

var NET = NETWORK{
	host:      "http://localhost",
	method:    0,
	parameter: "fuzzer",
	crypt:     false,
	status:    false,
}

type config struct {
	url    string
	method string
	form   *bytes.Buffer
	jar    []string
	proxy  string
	cmd    string
}

type NETWORK struct {
	host      string
	method    int
	parameter string
	crypt     bool
	status    bool

	_config       config
	_lastRequest  http.Request
	_lastResponse *http.Response
}

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) Setup(c *cli.Context) {
	url := c.String("u")
	method := c.Int("m")
	parameter := c.String("p")
	crypt := false

	if url == "" {
		fmt.Println("flag -u (url) is required")
		return
	}

	if parameter == "" {
		parameter = n.parameter
	}

	n.host = url
	n.method = method
	n.parameter = parameter
	n.crypt = crypt

	n.status = true
}

func (n *NETWORK) _getHandleBack() string {
	var r string

	if n.method == 0 || n.method == 1 {
		r = "echo($r);exit();"
	} else if n.method == 2 {
		r = "header('" + n.parameter + ":' . $r);exit();"
	} else if n.method == 3 {
		r = "setcookie('" + n.parameter + "', $r);exit();"
	}

	return r
}

func (n *NETWORK) _initConfig(r string) {
	c := config{
		url:    n.host,
		method: "GET",
	}

	n.status = true

	request := r + n._getHandleBack()

	if n.method == 0 { //GET
		c.url = n.host + "?" + n.parameter + "=" + request
	} else if n.method == 1 { //POST
		c.method = "POST"

		form := url.Values{}
		form.Set(n.parameter, request)
		c.form = bytes.NewBufferString(form.Encode())
	} else if n.method == 3 { //COOKIE
		c.jar = []string{}
	} else {
		n.status = false
	}

	n._config = c
}

func (n *NETWORK) _headerConfig(req *http.Request) {
	if n.method == 1 {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	} else if n.method == 2 {
		req.Header.Add(n.parameter, n._config.cmd)
	}
}

func (n *NETWORK) _handleSuccess(response *http.Response) {
	n._lastResponse = response
}

func (n *NETWORK) Send(r string) {
	client := &http.Client{}
	n._initConfig(r)

	if n.status == false {
		err := new(error)
		panic(err)
	}

	req, err := http.NewRequest(n._config.method, n._config.url, n._config.form)

	if err != nil {
		panic(err)
	}

	n._lastRequest = *req
	fmt.Println(n._lastRequest)

	n._headerConfig(req)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	n._handleSuccess(resp)
}

func (n *NETWORK) RequestInfo(c *cli.Context) {
	if n._lastResponse == nil {
		fmt.Println("You havn't made a request. You must make a request before seeing any information")
		return
	}

	if c.Bool("url") {
		showRequestUrl(&n._lastRequest)
	}

	if c.Bool("method") {
		showRequestMethod(&n._lastRequest)
	}

	if c.Bool("body") {
		showRequestBody(&n._lastRequest)
	}

	if c.Bool("headers") {
		showRequestHeaders(&n._lastRequest)
	}
}

func (n *NETWORK) ResponseInfo(c *cli.Context) {
	if n._lastResponse == nil {
		fmt.Println("You havn't made a request. You must make a request before seeing any information")
		return
	}

	if c.Bool("status") {
		showResponseStatus(n._lastResponse)
	}

	if c.Bool("body") {
		showResponseBody(n._lastResponse)
	}

	if c.Bool("headers") {
		showResponseHeaders(n._lastResponse)
	}

	if c.Bool("request") {
		showResponseRequest(n._lastResponse)
	}
}
