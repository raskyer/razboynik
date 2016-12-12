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

	action = phpadapter.CreateCMD(kc.GetRaw(), c.Shellscope, c.Shellmethod) + phpadapter.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, c)

	_, err = kc.Send(request)

	return kc, err
}
