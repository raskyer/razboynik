package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Listcmd struct{}

func (list *Listcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		action   string
		scope    string
		args     []string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	scope = "__DIR__"
	args = kl.GetArg()

	if config.Shellscope != "" {
		scope = "'" + config.Shellscope + "'"
	}

	if len(args) > 0 {
		scope = "'" + args[0] + "'"
	}

	action = "$r=implode('\n', scandir(" + scope + "));" + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return kernel.KernelResponse{Body: response.GetResult()}
}

func (list *Listcmd) GetName() string {
	return "-list"
}

func (list *Listcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
