package kernelmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/sysgo"
)

type Syscmd struct{}

func (sys *Syscmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		action string
		result string
		err    error
	)

	action = strings.Join(kl.GetArr(), " ")
	result, err = sysgo.Call(action)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), result)

	return kernel.KernelResponse{Err: err, Body: result}
}

func (sys *Syscmd) GetName() string {
	return "-sys"
}

func (sys *Syscmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
