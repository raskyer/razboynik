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

	if req.c.Method != M_GET && req.c.Method != M_POST && req.c.Method != M_HEADER && req.c.Method != M_COOKIE {
		req.c.Method = M_GET
	}

	if req.c.Parameter == "" {
		req.c.Parameter = PARAMETER
	}

	return nil
}
