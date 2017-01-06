package shellmodule

import (
	"net/http/httputil"

	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Shcmd struct {
	Debug bool
}

func (sh *Shcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		a, raw   string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	raw = strings.TrimSuffix(kl.GetRaw(), "--debug")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	if sh.Debug {
		kernel.WriteSuccess(kl.GetStdout(), "- REQUEST")
		b, _ := httputil.DumpRequestOut(request.GetHTTP(), true)
		kernel.WriteSuccess(kl.GetStdout(), string(b))

		kernel.WriteSuccess(kl.GetStdout(), "\n")
		kernel.WriteSuccess(kl.GetStdout(), "- RESPONSE\n\n")
		b, _ = httputil.DumpResponse(response.GetHTTP(), true)
		kernel.WriteSuccess(kl.GetStdout(), string(b))
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return kernel.KernelResponse{Err: err}
}

func (sh *Shcmd) GetName() string {
	return "sh"
}

func (sh *Shcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, true
}
