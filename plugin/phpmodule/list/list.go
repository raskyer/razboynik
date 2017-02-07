package phpmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Listitem = kernel.Item{
	Name: "-list",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			scope    string
			args     []string
			err      error
			response *razboy.RESPONSE
		)

		args = l.GetArg()
		scope = getScope(args, config)
		response, err = List(scope, config)

		if err != nil {
			return kernel.Response{Err: err}
		}

		kernel.WriteSuccess(l.GetStdout(), response.GetResult())

		return kernel.Response{Body: response.GetResult()}
	},
}

func List(scope string, config *razboy.Config) (*razboy.RESPONSE, error) {
	var (
		action  string
		request *razboy.REQUEST
	)

	action = "$r=implode('\n', scandir(" + scope + "));" + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)

	return razboy.Send(request)
}

func getScope(args []string, config *razboy.Config) string {
	scope := "__DIR__"

	if config.Shellscope != "" {
		scope = "'" + config.Shellscope + "'"
	}

	if len(args) > 0 {
		scope = "'" + args[0] + "'"
	}

	return scope
}
