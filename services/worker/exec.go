package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		k  *kernel.Kernel
		kc *kernel.KernelCmd
	)

	kc = kernel.CreateCmd(cmd)

	k = kernel.Boot()

	return k.Exec(kc, request)
}
