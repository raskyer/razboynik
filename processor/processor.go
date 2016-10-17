package processor

import (
	"net/http"

	"github.com/eatbytes/fuzz/network"
)

func Send(srv *network.NETWORK, str string) (*http.Response, error) {
	req, err := srv.Prepare(str)

	if err != nil {
		return nil, err
	}

	resp, err := srv.Send(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Process(srv *network.NETWORK, str string) (string, error) {
	resp, err := Send(srv, str)

	if err != nil {
		return "", err
	}

	result := srv.GetResultStr(resp)

	return result, nil
}
