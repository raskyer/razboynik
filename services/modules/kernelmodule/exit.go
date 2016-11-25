package kernelmodule

import (
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exit(kc *kernel.KernelCmd, config *config.Config) (*kernel.KernelCmd, error) {
	kernel.Boot().Stop()

	return kc, nil
}
