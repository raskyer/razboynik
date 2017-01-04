package kernel

import (
	"github.com/eatbytes/razboy"
	"github.com/natefinch/pie"
)

type KernelExternalArgs struct {
	Line   string
	Config *razboy.Config
}

type KernelExternalCmd interface {
	Exec(KernelExternalArgs, *string) error
}

func CreateProvider(kc KernelExternalCmd) (*pie.Server, error) {
	var (
		provider pie.Server
		err      error
	)

	provider = pie.NewProvider()
	err = provider.RegisterName("Plugin", kc)

	return &provider, err
}
