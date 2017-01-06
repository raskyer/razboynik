package phpmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Phpcmd struct{}

func (php *Phpcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		action   string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	action = strings.Join(kl.GetArr(), " ")
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return kernel.KernelResponse{Body: response.GetResult()}
}

func (php *Phpcmd) GetName() string {
	return "-php"
}

func (php *Phpcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
