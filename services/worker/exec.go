package worker

import (
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, config *config.Config) (*kernel.KernelCmd, error) {
	var (
		k   *kernel.Kernel
		kc  *kernel.KernelCmd
		err error
	)

	kc = kernel.CreateCmd(cmd)
	k = kernel.Boot()

	kc, err = k.Exec(kc, config)

	return kc, err
}
