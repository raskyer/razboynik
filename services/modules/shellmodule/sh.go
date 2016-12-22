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

func (sh *Shcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		a, raw   string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	sh.InitFlags(kl.GetArr())

	raw = strings.TrimSuffix(kl.GetRaw(), "--debug")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.CreateAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		kl.WriteError(err)

		return sh, err
	}

	if sh.Debug {
		kl.WriteSuccess("- REQUEST")
		b, _ := httputil.DumpRequestOut(request.GetHTTP(), true)
		kl.WriteSuccess(string(b))

		kl.WriteSuccess("\n")
		kl.WriteSuccess("- RESPONSE\n\n")
		b, _ = httputil.DumpResponse(response.GetHTTP(), true)
		kl.WriteSuccess(string(b))
	}

	kl.WriteSuccess(response.GetResult())

	return sh, err
}

func (sh *Shcmd) GetName() string {
	return "sh"
}

func (sh *Shcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, true
}

func (sh *Shcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (sh *Shcmd) GetResultStr() string {
	return ""
}
