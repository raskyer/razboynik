package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Exitcmd struct{}

func (exit *Exitcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	kernel.Boot().StopRun()

	return nil
}

func (e *Exitcmd) GetName() string {
	return "exit"
}

func (e *Exitcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
