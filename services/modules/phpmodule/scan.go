package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Scan(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		action  string
		err     error
		request *razboy.REQUEST
	)

	action = razboy.CreateScan() + razboy.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, c)

	_, err = kc.Send(request)

	return kc, err
}
