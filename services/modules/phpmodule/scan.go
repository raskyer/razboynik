package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Scancmd struct {
	result string
}

func (scan *Scancmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		action   string
		result   string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	action = razboy.CreateScan() + razboy.CreateAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return scan, err
	}

	result = response.GetResult()
	kl.WriteSuccess(result)

	return scan, nil
}

func (scan *Scancmd) GetName() string {
	return "-scan"
}

func (scan *Scancmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (scan *Scancmd) GetResult() []byte {
	return make([]byte, 0)
}

func (scan *Scancmd) GetResultStr() string {
	return ""
}
