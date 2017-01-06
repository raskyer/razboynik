package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Cdcmd struct{}

func (cd *Cdcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		raw, a, scope string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	raw = "cd " + strings.Join(kl.GetArr(), " ")

	if strings.Contains(raw, "&&") || strings.Contains(raw, "-") {
		return kernel.KernelResponse{}
	}

	raw += " && pwd"

	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(a, config)

	response, err = razboy.Send(request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope != "" {
		config.Shellscope = scope
		kernel.Boot().UpdatePrompt(config.Url, scope)
	}

	return kernel.KernelResponse{}
}

func (cd *Cdcmd) GetName() string {
	return "cd"
}

func (cd *Cdcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
