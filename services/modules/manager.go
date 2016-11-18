package modules

import (
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/shellmodule"
)

func Boot() {
	var k *kernel.Kernel

	k = kernel.Boot(kernel.KernelItem{
		Name: "raw",
		Fn:   shellmodule.Raw,
	})

	k.AddItem(kernel.KernelItem{
		Name: "cd",
		Fn:   shellmodule.Cd,
	})
}
