package shellmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Shcmd struct{}

func (sh *Shcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		a        string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	a = razboy.CreateCMD(kl.GetRaw(), config.Shellscope, config.Shellmethod) + razboy.CreateAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		kl.WriteError(err)

		return sh, err
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
