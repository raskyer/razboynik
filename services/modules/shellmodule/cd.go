package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Cd(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		action  string
		scope   string
		err     error
	)

	if strings.Contains(kc.GetRaw(), "&&") || kc.GetArrItem(1, "") == "-" {
		return Raw(kc, c)
	}

	action = phpadapter.CreateCD(kc.GetRaw(), kc.GetScope(), c.Shellmethod) + phpadapter.CreateAnswer(c.Method, c.Parameter)
	request = razboy.CreateRequest(action, kc.GetScope(), c)

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
