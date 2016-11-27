package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func ReadFile(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		action  string
		file    string
		err     error
	)

	file = kc.GetArrItem(1)

	if file == "" {
		return kc, errors.New("You should give the path of the file")
	}

	action = "$r=file_get_contents('" + file + "');" + phpadapter.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, kc.GetScope(), c)

	_, err = kc.Send(request)

	return kc, err
}
