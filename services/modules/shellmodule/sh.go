package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Shitem = kernel.Item{
	Name: "vim",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			a, raw   string
			err      error
			request  *razboy.REQUEST
			response *razboy.RESPONSE
		)

		raw = strings.TrimSuffix(l.GetRaw(), "--debug")
		a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

		request = razboy.CreateRequest(a, config)
		response, err = razboy.Send(request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		// if sh.Debug {
		// 	kernel.WriteSuccess(kl.GetStdout(), "- REQUEST")
		// 	b, _ := httputil.DumpRequestOut(request.GetHTTP(), true)
		// 	kernel.WriteSuccess(kl.GetStdout(), string(b))

		// 	kernel.WriteSuccess(kl.GetStdout(), "\n")
		// 	kernel.WriteSuccess(kl.GetStdout(), "- RESPONSE\n\n")
		// 	b, _ = httputil.DumpResponse(response.GetHTTP(), true)
		// 	kernel.WriteSuccess(kl.GetStdout(), string(b))
		// }

		kernel.WriteSuccess(l.GetStdout(), response.GetResult())

		return kernel.Response{Err: err}
	},
}
