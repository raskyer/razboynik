package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Listcmd struct {
	result string
}

func (list *Listcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		action   string
		scope    string
		args     []string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	scope = "__DIR__"
	args = kl.GetArr()

	if config.Shellscope != "" {
		scope = "'" + config.Shellscope + "'"
	}

	if len(args) > 0 {
		scope = "'" + args[0] + "'"
	}

	action = "$r=implode('\n', scandir(" + scope + "));" + razboy.CreateAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return list, err
	}

	list.result = response.GetResult()
	kl.WriteSuccess(list.result)

	return list, nil
}

func (list *Listcmd) GetName() string {
	return "-list"
}

func (list *Listcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (list *Listcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (list *Listcmd) GetResultStr() string {
	return list.result
}
