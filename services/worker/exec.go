package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, config *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		k  *kernel.Kernel
		kc *kernel.KernelCmd
	)

	k = kernel.Boot()
	kc = kernel.CreateCmd(cmd)

	return k.Exec(kc, config)
}
