package shellmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Raw(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		err     error
	)

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

	request.Action = phpadapter.CreateCMD(kc.GetRaw(), kc.GetScope(), c.Shellmethod, true, c.Method, c.Parameter)

	_, err = kc.Send(request)

	return kc, err
}
