package kernelmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/sysgo"
)

var Sysitem = kernel.Item{
	Name: "-sys",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			action string
			result string
			err    error
		)

		action = strings.Join(l.GetArg(), " ")
		result, err = sysgo.Call(action)

		if err != nil {
			return kernel.Response{Err: err}
		}

		kernel.WriteSuccess(l.GetStdout(), result)

		return kernel.Response{Err: err, Body: result}
	},
}
