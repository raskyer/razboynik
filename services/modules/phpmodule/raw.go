package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Raw(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		rzRes *razboy.RazResponse
		err   error
	)

	request.Type = "PHP"
	request.Action = kc.GetStr()

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	return kc, err
}
