package fuzzer

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

var NET = NETWORK{
	host:      "",
	method:    0,
	parameter: "Fuzzer",
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

	req, err := http.NewRequest(c.method, c.url, bytes)

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

		return req, false
	}

	if n.method == 1 {
		config = n.postConfig(r)
		req, err = http.NewRequest(config.method, config.url, config.form)

		if err != nil {
			fmt.Println("Error: Impossible to create request with config :")
			fmt.Println(config)
			return nil, true
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		return req, false
	}

	if n.method == 2 {
		config = n.headerConfig(r)
		req, err = http.NewRequest(config.method, config.url, nil)

		if err != nil {
			fmt.Println("Error: Impossible to create request with config :")
			fmt.Println(config)
			return nil, true
		}

		req.Header.Add(n.parameter, n.cmd)

		return req, false
	}

	if n.method == 3 {
	}

	return req, true
}

func (n *NETWORK) Send(req *http.Request) (*http.Response, bool) {
	if n.status != true {
		fmt.Println("You havn't setup the required information, please refer to srv config")
		return nil, true
	}

	n._respBody = nil

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error: Impossible to send request. Error msg : ")
		fmt.Println(err)
		fmt.Println(resp.Status)
		return nil, true
	}

	n._lastResponse = resp

	return resp, false
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

func (n *NETWORK) headerConfig(r string) *config {
	n.status = true

	request := Encode(r)
	n.cmd = request

	c := config{
		url:    n.host,
		method: "GET",
		form:   nil,
	}

	return &c
}

func (n *NETWORK) cookieConfig(r string) {
}
