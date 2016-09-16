package fuzzer

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

type callback func(string)

var NET = NETWORK{
	host:      "",
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
	file   bool
}

type NETWORK struct {
	host      string
	method    int
	parameter string
	crypt     bool
	status    bool
	cmd       string

	_body         url.Values
	_respBody     []byte
	_lastResponse *http.Response
}

func (n *NETWORK) GetUrl() string {
	return n.host
}

func (n *NETWORK) GetMethod() int {
	return n.method
}

func (n *NETWORK) GetMethodStr() string {
	if n.method == 0 {
		return "GET"
	}

	if n.method == 1 {
		return "POST"
	}

	if n.method == 3 {
		return "HEADER"
	}

	if n.method == 4 {
		return "COOKIE"
	}

	return "ERROR"
}

func (n *NETWORK) GetParameter() string {
	return n.parameter
}

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) SetConfig(url string, method int, parameter string, crypt bool) {
	n.host = url
	n.method = method
	n.parameter = parameter
	n.crypt = crypt

	n.status = true
}

func (n *NETWORK) PrepareUpload(bytes *bytes.Buffer, bondary string) (*http.Request, bool) {
	if n.status != true {
		fmt.Println("You havn't setup the required information, please refer to srv config")
		return nil, true
	}

	n.status = true

	c := config{
		url:    n.host,
		method: "POST",
		form:   bytes,
	}

	req, err := http.NewRequest("POST", n.host, bytes)

	req.Header.Set("Content-Type", bondary)

	if err != nil {
		fmt.Println("Error: Impossible to create request with config :")
		fmt.Println(c)
		return nil, true
	}

	return req, false
}

func (n *NETWORK) Prepare(r string) (*http.Request, bool) {
	if n.status != true {
		fmt.Println("You havn't setup the required information, please refer to srv config")
		return nil, true
	}

	var config *config
	var req *http.Request
	var err error

	if n.method == 0 {
		config = n.getConfig(r)
		req, err = http.NewRequest(config.method, config.url, nil)

		if err != nil {
			fmt.Println("Error: Impossible to create request with config :")
			fmt.Println(config)
			return nil, true
		}
	}

	if n.method == 1 {
		config = n.postConfig(r)
		req, err = http.NewRequest(config.method, config.url, config.form)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			fmt.Println("Error: Impossible to create request with config :")
			fmt.Println(config)
			return nil, true
		}
	}

	if n.method == 3 {
	}

	if n.method == 4 {
	}

	return req, false
}

func (n *NETWORK) Send(req *http.Request, f callback) {
	if n.status != true {
		fmt.Println("You havn't setup the required information, please refer to srv config")
		return
	}

	var httpResponse *http.Response
	var err int
	var response string

	n._respBody = nil

	httpResponse, err = n._send(req)

	if err == 1 {
		return
	}

	buffer := GetBody(httpResponse)
	response = string(buffer)

	if httpResponse != nil && httpResponse.StatusCode < 400 {
		f(response)
	} else {
		fmt.Println("Error with the response: " + httpResponse.Status)
	}
}

func (n *NETWORK) postConfig(r string) *config {
	n.status = true

	request := Encode(r)
	n.cmd = request

	form := url.Values{}
	form.Set(n.parameter, request)
	n._body = form

	data := bytes.NewBufferString(form.Encode())

	c := config{
		url:    n.host,
		method: "POST",
		form:   data,
	}

	return &c
}

func (n *NETWORK) getConfig(r string) *config {
	n.status = true

	request := Encode(r)
	n.cmd = request

	url := n.host + "?" + n.parameter + "=" + request

	c := config{
		url:    url,
		method: "GET",
		form:   nil,
	}

	return &c
}

func (n *NETWORK) headerConfig(r string) {
}

func (n *NETWORK) cookieConfig(r string) {
}

func (n *NETWORK) _send(req *http.Request) (*http.Response, int) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error: Impossible to send request. Error msg : ")
		fmt.Println(err)
		return nil, 1
	}

	n._lastResponse = resp

	return resp, 0
}

func (n *NETWORK) GetResponse() *http.Response {
	return n._lastResponse
}

func (n *NETWORK) GetRequest() *http.Request {
	n._lastResponse.Request.PostForm = n._body
	return n._lastResponse.Request
}
