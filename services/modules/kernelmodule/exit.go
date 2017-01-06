package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func CreateExit() kernel.KernelCommand {
	return new(Exitcmd)
}

type Exitcmd struct{}

func (exit *Exitcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	kernel.Boot().StopRun()
	return kernel.KernelResponse{}
}

func (e *Exitcmd) GetName() string {
	return "exit"
}

func (e *Exitcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
