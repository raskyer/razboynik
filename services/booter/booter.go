package booter

import (
	"fmt"

	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/examplemodule"
	"github.com/eatbytes/razboynik/services/modules/kernelmodule"
	"github.com/eatbytes/razboynik/services/modules/phpmodule"
	"github.com/eatbytes/razboynik/services/modules/shellmodule"
	"github.com/eatbytes/razboynik/services/worker/targetwork"
)

func Boot() {
	var (
		k        *kernel.Kernel
		items    []*kernel.Item
		rpcitems []*kernel.Item
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
	rpcitems = bootRPC()
	items = append(items, rpcitems...)

	k.SetDefault(&shellmodule.Shitem)
	k.SetCommands(items)
}

func bootRPC() []*kernel.Item {
	var (
		config *targetwork.Configuration
		err    error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		fmt.Println(err)
		return []*kernel.Item{}
	}

	return config.Providers
}
