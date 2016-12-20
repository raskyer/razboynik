package razboy

import (
	"errors"
	"net/http"
)

const KEY = "RAZBOYNIK_KEY"
const PARAM = "razboynik"

func main() {}

func Send(req *REQUEST) (*RESPONSE, error) {
	var (
		res *RESPONSE
		err error
	)

	err = Check(req)

	if err != nil {
		return nil, err
	}

	err = Prepare(req)

	if err != nil {
		return nil, err
	}

	res, err = SendRequest(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func Prepare(req *REQUEST) error {
	var err error

	err = Check(req)

	if err != nil {
		return err
	}

	if req.Upload {
		err = _createUploadRequest(req)
	} else {
		err = _createSimpleRequest(req)
	}

	return err
}

func SendRequest(req *REQUEST) (*RESPONSE, error) {
	var (
		res    *RESPONSE
		client *http.Client
		resp   *http.Response
		err    error
	)

	if !req.setup {
		return nil, errors.New("Problem with request")
	}

	client = &http.Client{}
	resp, err = client.Do(req.http)

	if err != nil {
		return nil, err
	}

	res = &RESPONSE{
		http:    resp,
		request: req,
	}

	return res, nil
}
