package phpmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Phpcmd struct{}

func (php *Phpcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		action   string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	action = strings.Join(kl.GetArr(), " ")
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return php, err
	}

	kl.WriteSuccess(response.GetResult())

	return php, nil
}

func (php *Phpcmd) GetName() string {
	return "-php"
}

func (php *Phpcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (php *Phpcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (php *Phpcmd) GetResultStr() string {
	return ""
}
