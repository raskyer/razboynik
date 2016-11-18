package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(config *core.REQUEST) (kernel.KernelCmd, error) {
	var (
		kc kernel.KernelCmd
		fn kernel.KernelFunction
	)

	kc = kernel.CreateCmd(config.SHLc.Cmd, config.SHLc.Scope)
	fn = _exec_getFunction(kc.GetName())

	return fn(kc, config)
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
