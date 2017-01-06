package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Deletecmd struct{}

func (delete *Deletecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		action  string
		scope   string
		args    []string
		err     error
		request *razboy.REQUEST
	)

	args = kl.GetArg()

	if len(args) < 1 {
		return kernel.KernelResponse{Err: errors.New("You should give the path of the file to delete")}
	}

	scope = args[0]

	if config.Shellscope != "" {
		scope = config.Shellscope + "/" + scope
	}

	action = "if(is_dir('" + scope + "')){$r=rmdir('" + scope + "');}else{$r=unlink('" + scope + "');}" + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)

	_, err = razboy.Send(request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), "Delete successfully")

	return kernel.KernelResponse{Body: true}
}

func (delete *Deletecmd) GetName() string {
	return "-delete"
}

func (delete *Deletecmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
