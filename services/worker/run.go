package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Run(request *core.REQUEST) error {
	var (
		k *kernel.Kernel
	)

	k = kernel.Boot()
	k.Run(request)

	return nil
}
