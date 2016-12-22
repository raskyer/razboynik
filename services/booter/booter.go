package booter

import (
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/shellmodule"
)

func Boot() {
	var (
		k *kernel.Kernel
	)

	k = kernel.Boot()
	k.SetDefault(new(shellmodule.SHCmd))

	k.SetCommands([]kernel.KernelCommand{
		new(shellmodule.HelloWorldCmd),
		new(shellmodule.CDCmd),
		new(shellmodule.PWDCmd),
	})
}
