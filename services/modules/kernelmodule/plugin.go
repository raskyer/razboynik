package kernelmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/lister"
)

type Plugincmd struct{}

func (plugin *Plugincmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		err      error
		args     *kernel.KernelExternalArgs
		response *kernel.KernelExternalResponse
	)

	if len(kl.GetArg()) < 1 {
		return errors.New("Please specify a path for external plugin")
	}

	args = new(kernel.KernelExternalArgs)
	args.Line = kl.GetStr()

	response, err = kernel.ExecuteProvider(kl.GetArg()[0], "Exec", args)

	if err != nil {
		return err
	}

	kernel.WriteSuccess(kl.GetStdout(), response.Response)

	return nil
}

func (plugin *Plugincmd) GetName() string {
	return "-plugin"
}

func (plugin *Plugincmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return lister.Local, true
}
