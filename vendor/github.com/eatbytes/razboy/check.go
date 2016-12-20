package razboy

import (
	"errors"
)

func Check(req *REQUEST) error {
	if req == nil {
		return errors.New("Empty pointer")
	}

	return _checkSERVER(req)
}

func _checkSERVER(req *REQUEST) error {
	if req.c.Url == "" {
		return errors.New("REQUEST [url] should not be empty")
	}

	if req.c.Method != "GET" && req.c.Method != "POST" && req.c.Method != "HEADER" && req.c.Method != "COOKIE" {
		req.c.Method = "GET"
	}

	if req.c.Parameter == "" {
		req.c.Parameter = PARAM
	}

	return nil
}
