package booter

import (
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/kernelmodule"
	"github.com/eatbytes/razboynik/services/modules/shellmodule"
)

func Boot() {
	var (
		k *kernel.Kernel
		i []*kernel.KernelItem
	)

	k = kernel.Boot(&kernel.KernelItem{
		Name: "raw",
		Fn:   shellmodule.Raw,
	})

	i = []*kernel.KernelItem{
		&kernel.KernelItem{
			Name: "cd",
			Fn:   shellmodule.Cd,
		},
		&kernel.KernelItem{
			Name: "exit",
			Fn:   kernelmodule.Exit,
		},
	}

	k.SetItems(i)
}
