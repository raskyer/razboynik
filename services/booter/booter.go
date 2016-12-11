package booter

import (
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/kernelmodule"
	"github.com/eatbytes/razboynik/services/modules/phpmodule"
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
			Name: "vim",
			Fn:   shellmodule.Vim,
		},
		&kernel.KernelItem{
			Name: "pwd",
			Fn:   shellmodule.Pwd,
		},
		&kernel.KernelItem{
			Name: "-php",
			Fn:   phpmodule.Raw,
		},
		&kernel.KernelItem{
			Name: "-listfile",
			Fn:   phpmodule.ListFile,
		},
		&kernel.KernelItem{
			Name: "-readfile",
			Fn:   phpmodule.ReadFile,
		},
		&kernel.KernelItem{
			Name: "-delete",
			Fn:   phpmodule.Delete,
		},
		&kernel.KernelItem{
			Name: "-upload",
			Fn:   phpmodule.Upload,
		},
		&kernel.KernelItem{
			Name: "-download",
			Fn:   phpmodule.Download,
		},
		&kernel.KernelItem{
			Name: "-sys",
			Fn:   kernelmodule.Sys,
		},
		&kernel.KernelItem{
			Name: "-debug",
			Fn:   kernelmodule.Debug,
		},
		&kernel.KernelItem{
			Name: "-encode",
			Fn:   kernelmodule.Encode,
		},
		&kernel.KernelItem{
			Name: "-decode",
			Fn:   kernelmodule.Decode,
		},
		&kernel.KernelItem{
			Name: "exit",
			Fn:   kernelmodule.Exit,
		},
	}

	k.SetItems(i)
}
