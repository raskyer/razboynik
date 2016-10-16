package networking

import (
	"net/http"

	"github.com/eatbytes/fuzzcore"
)

func Send(str string) (*http.Response, error) {
	req, err := fuzzcore.NET.Prepare(str)

	if err != nil {
		return nil, err
	}

	resp, err := fuzzcore.NET.Send(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Process(str string) (string, error) {
	resp, err := Send(str)

	if err != nil {
		return "", err
	}

	result := fuzzcore.NET.GetResultStr(resp)

	return result, nil
}
