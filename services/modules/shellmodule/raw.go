package shellmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Raw(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		action  string
		err     error
	)

	action = phpadapter.CreateCMD(kc.GetRaw(), kc.GetScope(), c.Shellmethod) + phpadapter.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, kc.GetScope(), c)

	_, err = kc.Send(request)

	return kc, err
}
