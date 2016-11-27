package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exit(kc *kernel.KernelCmd, config *razboy.Config) (*kernel.KernelCmd, error) {
	kernel.Boot().Stop()

	return kc, nil
}