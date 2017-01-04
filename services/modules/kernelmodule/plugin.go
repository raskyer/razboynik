package kernelmodule

import (
	"net/rpc"
	"os"
	"runtime"

	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/lister"
	"github.com/natefinch/pie"
)

type Plugincmd struct{}

func (plugin *Plugincmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		response string
		path     string
		err      error
		args     *kernel.KernelExternalArgs
		client   *rpc.Client
	)

	if len(kl.GetArg()) < 1 {
		return errors.New("Please specify a path for external plugin")
	}

	path = kl.GetArg()[0]

	if runtime.GOOS == "windows" {
		path = path + ".exe"
	}

	client, err = pie.StartProvider(os.Stderr, path)

	if err != nil {
		return err
	}

	defer client.Close()

	args = new(kernel.KernelExternalArgs)
	args.Line = kl.GetStr()

	err = client.Call("Plugin.Exec", args, &response)

	if err != nil {
		return err
	}

	kernel.WriteSuccess(kl.GetStdout(), response)

	return nil
}

func (plugin *Plugincmd) GetName() string {
	return "-plugin"
}

func (plugin *Plugincmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return lister.Local, true
}
