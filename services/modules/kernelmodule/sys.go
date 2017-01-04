package kernelmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/sysgo"
)

type Syscmd struct{}

func (sys *Syscmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		action string
		result string
		err    error
	)

	action = strings.Join(kl.GetArr(), " ")
	result, err = sysgo.Call(action)

	if err != nil {
		return err
	}

	kernel.WriteSuccess(kl.GetStdout(), result)

	return err
}

func (sys *Syscmd) GetName() string {
	return "-sys"
}

func (sys *Syscmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}

func (sys *Syscmd) GetResult() []byte {
	return make([]byte, 0)
}

func (sys *Syscmd) GetResultStr() string {
	return ""
}
