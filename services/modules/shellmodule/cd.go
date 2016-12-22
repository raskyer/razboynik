package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type CDCmd struct {
	Stdout string
	Stderr string
}

func (cd *CDCmd) Init(stdout, stderr string) {
	cd.Stdout = stdout
	cd.Stderr = stderr
}

func (cd *CDCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, int, error) {
	var (
		raw, a, scope string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	raw = "cd " + strings.Join(kl.GetArr(), " ")

	if strings.Contains(raw, "&&") || strings.Contains(raw, "-") {
		return nil, 0, nil
	}

	raw += " && pwd"

	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.CreateAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(a, config)

	response, err = razboy.Send(request)

	if err != nil {
		cd.WriteError(err)

		return cd, 1, err
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope != "" {
		config.Shellscope = scope
		kernel.Boot().UpdatePrompt(config.Url, scope)
	}

	return cd, 0, nil
}

func (cd *CDCmd) Write(e error, i ...interface{}) error {
	if e != nil {
		return cd.WriteError(e)
	}

	return cd.WriteSuccess(i...)
}

func (cd *CDCmd) WriteSuccess(i ...interface{}) error {
	return kernel.Boot().WriteSuccess(cd.Stdout, i...)
}

func (cd *CDCmd) WriteError(e error) error {
	return kernel.Boot().WriteError(cd.Stderr, e)
}

func (cd *CDCmd) GetName() string {
	return "cd"
}

func (cd *CDCmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, true
}

func (cd *CDCmd) GetResult() []byte {
	return make([]byte, 0)
}

func (cd *CDCmd) GetResultStr() string {
	return ""
}
