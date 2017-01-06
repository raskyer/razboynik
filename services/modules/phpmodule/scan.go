package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Scancmd struct {
	result string
}

func (scan *Scancmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
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
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return kernel.KernelResponse{Body: response.GetResult()}
}

func (scan *Scancmd) GetName() string {
	return "-scan"
}

func (scan *Scancmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
