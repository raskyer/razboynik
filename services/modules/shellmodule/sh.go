package shellmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type SHCmd struct {
	Stdout string
	Stderr string
}

func (sh *SHCmd) Init(stdout, stderr string) {
	sh.Stdout = stdout
	sh.Stderr = stderr
}

func (sh *SHCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, int, error) {
	var (
		a        string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	a = razboy.CreateCMD(kl.GetRaw(), config.Shellscope, config.Shellmethod) + razboy.CreateAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)
	err = sh.Write(err, response.GetResult())

	return sh, 0, err
}

func (sh *SHCmd) Write(e error, i ...interface{}) error {
	if e != nil {
		return sh.WriteError(e)
	}

	return sh.WriteSuccess(i...)
}

func (sh *SHCmd) WriteSuccess(i ...interface{}) error {
	return kernel.Boot().WriteSuccess(sh.Stdout, i...)
}

func (sh *SHCmd) WriteError(e error) error {
	return kernel.Boot().WriteError(sh.Stderr, e)
}

func (sh *SHCmd) GetName() string {
	return "sh"
}

func (sh *SHCmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, true
}

func (sh *SHCmd) GetResult() []byte {
	return make([]byte, 0)
}

func (sh *SHCmd) GetResultStr() string {
	return ""
}
