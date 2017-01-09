package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Exititem = kernel.Item{
	Name: "exit",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		k := kernel.Boot()
		k.StopRun()

		return kernel.Response{}
	},
}
