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
		k     *kernel.Kernel
		items []*kernel.Item
	)

	k = kernel.Boot()

	items = []*kernel.Item{
		&shellmodule.Cditem,
		&shellmodule.Pwditem,
		&shellmodule.Vimitem,
		&phpmodule.Phpitem,
		&phpmodule.Downloaditem,
		&phpmodule.Uploaditem,
		&phpmodule.Listitem,
		&phpmodule.Readfileitem,
		&phpmodule.Scanitem,
		&phpmodule.Deleteitem,
		&kernelmodule.Decodeitem,
		&kernelmodule.Encodeitem,
		&kernelmodule.Sysitem,
		&examplemodule.Fiboitem,
		&kernelmodule.Exititem,
	}

	k.SetDefault(&shellmodule.Shitem)
	k.SetCommands(items)
}
