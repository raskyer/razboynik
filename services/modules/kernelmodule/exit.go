package kernelmodule

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exit(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	kernel.Boot().Stop()

	return kc, nil
}
