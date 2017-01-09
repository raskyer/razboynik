package booter

import (
	"encoding/json"
	"os"

	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/examplemodule"
	"github.com/eatbytes/razboynik/services/modules/kernelmodule"
	"github.com/eatbytes/razboynik/services/modules/phpmodule"
	"github.com/eatbytes/razboynik/services/modules/shellmodule"
	"github.com/eatbytes/razboynik/services/usr"
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
		file          *os.File
		decoder       *json.Decoder
		items         []*kernel.Item
		filepath, dir string
		err           error
	)

	dir, err = usr.GetHomeDir()

	if err != nil {
		return items
	}

	filepath = dir + "/.razboynik.providers.json"

	if err != nil {
		return items
	}

	file, err = os.Open(filepath)
	defer file.Close()

	if err != nil {
		return items
	}

	decoder = json.NewDecoder(file)
	err = decoder.Decode(items)

	if err != nil {
		return items
	}

	return items
}
