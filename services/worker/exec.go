package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, config *razboy.Config) error {
	var k *kernel.Kernel

	k = kernel.Boot()

	return k.Exec(cmd, config)
}
