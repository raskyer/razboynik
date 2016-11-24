package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		k   *kernel.Kernel
		kc  *kernel.KernelCmd
		err error
	)

	kc = kernel.CreateCmd(cmd)
	k = kernel.Boot()

	kc, err = k.Exec(kc, request)

	return kc, err
}
