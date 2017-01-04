package shellmodule

import (
	"net/http/httputil"

	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/spf13/pflag"
)

type Shcmd struct {
	Debug bool
}

func (sh *Shcmd) InitFlags(args []string) {
	flaghandler := pflag.NewFlagSet("/bin/sh", pflag.ContinueOnError)
	flaghandler.BoolVar(&sh.Debug, "debug", false, "Debug mode")
	flaghandler.Parse(args)
}

func (sh *Shcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		a, raw   string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	//sh.InitFlags(kl.GetArr())

	raw = strings.TrimSuffix(kl.GetRaw(), "--debug")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		return err
	}

	if sh.Debug {
		kernel.WriteSuccess(kl.GetStdout(), "- REQUEST")
		b, _ := httputil.DumpRequestOut(request.GetHTTP(), true)
		kernel.WriteSuccess(kl.GetStdout(), string(b))

		kernel.WriteSuccess(kl.GetStdout(), "\n")
		kernel.WriteSuccess(kl.GetStdout(), "- RESPONSE\n\n")
		b, _ = httputil.DumpResponse(response.GetHTTP(), true)
		kernel.WriteSuccess(kl.GetStdout(), string(b))
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return err
}

func (sh *Shcmd) GetName() string {
	return "sh"
}

func (sh *Shcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, true
}
