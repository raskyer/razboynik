package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, config *razboy.Config) (*kernel.KernelCmd, error) {
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
