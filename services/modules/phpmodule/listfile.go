package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func ListFile(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		rzRes *razboy.RazResponse
		scope string
		err   error
	)

	scope = "__DIR__"

	if kc.GetScope() != "" {
		scope = "'" + kc.GetScope() + "'"
	}

	if kc.GetStr() != "" {
		scope = "'" + kc.GetStr() + "'"
	}

	request.Type = "PHP"
	request.Action = "$r=json_encode(scandir(" + scope + "));" + phpadapter.CreateAnswer(request)

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	return kc, err
}
