package worker

import (
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Run(config *config.Config) error {
	var (
		k *kernel.Kernel
	)

	k = kernel.Boot()
	k.Run(config)

	return nil
}
