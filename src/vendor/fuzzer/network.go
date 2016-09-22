package fuzzer

import (
	"bytes"
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

func (n *NETWORK) PrepareUpload(bytes *bytes.Buffer, bondary string) (*http.Request, error) {
	var ferr FuzzerError

	if n.status != true {
		ferr = SetupErr()
		return nil, ferr
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
		ferr := BuildRequestErr(err, &c)
		return nil, ferr
	}

	return req, nil
}

func (n *NETWORK) Prepare(r string) (*http.Request, error) {
	var ferr FuzzerError
	ferr = NoMethodFoundErr()

	if n.status != true {
		ferr = SetupErr()
		return nil, ferr
	}

	var config *config
	var req *http.Request
	var err error

	if n.method == 0 {
		config = n.getConfig(r)
		req, err = http.NewRequest(config.method, config.url, nil)

		if err != nil {
			ferr := BuildRequestErr(err, config)
			return nil, ferr
		}

		return req, nil
	}

	if n.method == 1 {
		config = n.postConfig(r)
		req, err = http.NewRequest(config.method, config.url, config.form)

		if err != nil {
			ferr := BuildRequestErr(err, config)
			return nil, ferr
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		return req, nil
	}

	if n.method == 2 {
		config = n.headerConfig(r)
		req, err = http.NewRequest(config.method, config.url, nil)

		if err != nil {
			ferr := BuildRequestErr(err, config)
			return nil, ferr
		}

		req.Header.Add(n.parameter, n.cmd)

		return req, nil
	}

	if n.method == 3 {
	}

	return req, ferr
}

func (n *NETWORK) Send(req *http.Request) (*http.Response, error) {
	var ferr FuzzerError

	if n.status != true {
		ferr = SetupErr()
		return nil, ferr
	}

	n._respBody = nil

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		ferr = RequestErr(err, resp.StatusCode)
		return nil, ferr
	}

	n._lastResponse = resp

	return resp, nil
}
