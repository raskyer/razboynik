package booter

import (
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/examplemodule"
	"github.com/eatbytes/razboynik/services/modules/kernelmodule"
	"github.com/eatbytes/razboynik/services/modules/phpmodule"
	"github.com/eatbytes/razboynik/services/modules/shellmodule"
)

func Boot() {
	var (
		k *kernel.Kernel
	)

	k = kernel.Boot()

	k.SetDefault(new(shellmodule.Shcmd))
	k.SetCommands([]kernel.KernelCommand{
		new(kernelmodule.Plugincmd),
		new(examplemodule.HelloWorldCmd),
		new(examplemodule.Fibocmd),
		new(shellmodule.Cdcmd),
		new(shellmodule.Pwdcmd),
		new(shellmodule.Vimcmd),
		new(phpmodule.Phpcmd),
		new(phpmodule.Downloadcmd),
		new(phpmodule.Uploadcmd),
		new(phpmodule.Listcmd),
		new(phpmodule.Readfilecmd),
		new(phpmodule.Scancmd),
		new(phpmodule.Deletecmd),
		new(kernelmodule.Decodecmd),
		new(kernelmodule.Encodecmd),
		new(kernelmodule.Syscmd),
		new(kernelmodule.Exitcmd),
	})
}
