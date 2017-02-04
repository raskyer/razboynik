package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Scanitem = kernel.Item{
	Name: "-scan",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			action   string
			err      error
			request  *razboy.REQUEST
			response *razboy.RESPONSE
		)

		action = razboy.CreateScan() + razboy.AddAnswer(config.Method, config.Parameter)
		request = razboy.CreateRequest(action, config)
		response, err = razboy.Send(request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		kernel.WriteSuccess(l.GetStdout(), response.GetResult())

		return kernel.Response{Body: response.GetResult()}
	},
}
