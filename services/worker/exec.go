package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Exec(cmd string, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var kc *kernel.KernelCmd
	kc = kernel.CreateCmd(cmd, request.SHLc.Scope)

	return kc.Exec(request)
}
