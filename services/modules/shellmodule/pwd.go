package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Pwdcmd struct{}

func (pwd *Pwdcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		raw, a, scope string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	raw = "pwd " + strings.Join(kl.GetArr(), " ")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		return err
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope != "" {
		config.Shellscope = scope
		kernel.Boot().UpdatePrompt(config.Url, scope)
	}

	return nil
}

func (pwd *Pwdcmd) GetName() string {
	return "pwd"
}

func (pwd *Pwdcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
