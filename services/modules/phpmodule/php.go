package phpmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Phpcmd struct{}

func (php *Phpcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
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
		return err
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return nil
}

func (php *Phpcmd) GetName() string {
	return "-php"
}

func (php *Phpcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
