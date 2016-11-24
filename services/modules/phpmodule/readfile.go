package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func ReadFile(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		rzRes *razboy.RazResponse
		file  string
		err   error
	)

	file = kc.GetArrItem(1)

	if file == "" {
		return kc, errors.New("You should give the path of the file")
	}

	request.Type = "PHP"
	request.Action = "$r=file_get_contents('" + file + "');" + phpadapter.CreateAnswer(request)

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	return kc, err
}
