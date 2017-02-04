package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Pwditem = kernel.Item{
	Name: "pwd",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			raw, a, scope string
			err           error
			request       *razboy.REQUEST
			response      *razboy.RESPONSE
		)

		raw = "pwd " + l.GetStr()
		a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

		request = razboy.CreateRequest(a, config)
		response, err = razboy.Send(request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		scope = strings.TrimSpace(response.GetResult())

		if scope != "" {
			config.Shellscope = scope
			kernel.Boot().UpdatePrompt(config.Url, scope)
		}

		return kernel.Response{Body: scope}
	},
}
