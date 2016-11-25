package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func ListFile(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
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

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

	request.Action = "$r=json_encode(scandir(" + scope + "));" + phpadapter.CreateAnswer(c.Method, c.Parameter)

	_, err = kc.Send(request)

	return kc, err
}
