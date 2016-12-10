package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func ListFile(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		action  string
		scope   string
		err     error
	)

	scope = "__DIR__"

	if kc.GetScope() != "" {
		scope = "'" + kc.GetScope() + "'"
	}

	if kc.GetStr() != "" {
		scope = "'" + kc.GetStr() + "'"
	}

	action = phpadapter.CreateListFile(scope) + phpadapter.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, kc.GetScope(), c)

	_, err = kc.Send(request)

	return kc, err
}
