package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Cdcmd struct{}

func (cd *Cdcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		raw, a, scope string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	raw = "cd " + strings.Join(kl.GetArr(), " ")

	if strings.Contains(raw, "&&") || strings.Contains(raw, "-") {
		return nil, nil
	}

	raw += " && pwd"

	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.CreateAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(a, config)

	response, err = razboy.Send(request)

	if err != nil {
		kl.WriteError(err)

		return cd, err
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope != "" {
		config.Shellscope = scope
		kernel.Boot().UpdatePrompt(config.Url, scope)
	}

	return cd, nil
}

func (cd *Cdcmd) GetName() string {
	return "cd"
}

func (cd *Cdcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (cd *Cdcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (cd *Cdcmd) GetResultStr() string {
	return ""
}
