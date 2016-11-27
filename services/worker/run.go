package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Run(config *razboy.Config) error {
	var (
		k *kernel.Kernel
	)

	k = kernel.Boot()
	k.Run(config)

	return nil
}
