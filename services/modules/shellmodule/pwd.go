package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type PWDCmd struct {
	Stdout string
	Stderr string
}

func (pwd *PWDCmd) Init(stdout, stderr string) {
	pwd.Stdout = stdout
	pwd.Stderr = stderr
}

func (pwd *PWDCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, int, error) {
	var (
		raw, a, scope string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	raw = "pwd " + strings.Join(kl.GetArr(), " ")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.CreateAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		pwd.WriteError(err)

		return pwd, 1, err
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope != "" {
		config.Shellscope = scope
		kernel.Boot().UpdatePrompt(config.Url, scope)
	}

	return pwd, 0, nil
}

func (pwd *PWDCmd) Write(e error, i ...interface{}) error {
	if e != nil {
		return pwd.WriteError(e)
	}

	return pwd.WriteSuccess(i...)
}

func (pwd *PWDCmd) WriteSuccess(i ...interface{}) error {
	return kernel.Boot().WriteSuccess(pwd.Stdout, i...)
}

func (pwd *PWDCmd) WriteError(e error) error {
	return kernel.Boot().WriteError(pwd.Stderr, e)
}

func (pwd *PWDCmd) GetName() string {
	return "pwd"
}

func (pwd *PWDCmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, true
}

func (pwd *PWDCmd) GetResult() []byte {
	return make([]byte, 0)
}

func (pwd *PWDCmd) GetResultStr() string {
	return ""
}
