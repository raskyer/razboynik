package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Delete(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		action  string
		scope   string
		err     error
	)

	scope = kc.GetArrItem(1)

	if scope == "" {
		return kc, errors.New("You should give the path of the file")
	}

	if kc.GetScope() != "" {
		scope = kc.GetScope() + "/" + scope
	}

	action = phpadapter.CreateDelete(scope) + phpadapter.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, kc.GetScope(), c)

	_, err = kc.Send(request)

	return kc, err
}
