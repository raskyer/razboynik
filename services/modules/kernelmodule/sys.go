package kernelmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/sysgo"
)

func Sys(kc *kernel.KernelCmd, config *razboy.Config) (*kernel.KernelCmd, error) {
	s, err := sysgo.Call(kc.GetStr())
	kc.SetResult(s)

	return kc, err
}
