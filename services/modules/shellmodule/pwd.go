package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Pwd(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		scope   string
		err     error
	)

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

	request.Action = phpadapter.CreateCMD(kc.GetRaw(), kc.GetScope(), c.Shellmethod, true, c.Method, c.Parameter)

	_, err = kc.Send(request)

	if err != nil {
		return kc, err
	}

	scope = strings.TrimSpace(kc.GetResult())

	if scope != "" {
		kc.SetScope(scope)
		kernel.Boot().UpdatePrompt(c.Url, scope)
	}

	return kc, nil
}
