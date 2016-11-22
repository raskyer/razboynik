package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		kc *kernel.KernelCmd
		fn kernel.KernelFunction
	)

	request.SHLc.Cmd = cmd

	kc = kernel.CreateCmd(request.SHLc.Cmd, request.SHLc.Scope)
	fn = _exec_getFunction(kc.GetName())

	return fn(kc, request)
}

func _exec_getFunction(name string) kernel.KernelFunction {
	var k *kernel.Kernel

	k = kernel.Boot()

	for _, item := range k.GetItems() {
		if item.Name == name {
			return item.Fn
		}
	}

	return k.GetDefaultItem().Fn
}
