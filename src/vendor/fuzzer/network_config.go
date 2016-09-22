package fuzzer

import (
	"bytes"
	"net/url"
)

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
