package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Exitcmd struct{}

func (exit *Exitcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	kernel.Boot().StopRun()

	return exit, nil
}

func (e *Exitcmd) GetName() string {
	return "exit"
}

func (e *Exitcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (e *Exitcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (e *Exitcmd) GetResultStr() string {
	return ""
}
