package phpmodule

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

// import (
// 	"errors"

// 	"github.com/eatbytes/razboy"
// 	"github.com/eatbytes/razboy/adapter/phpadapter"
// 	"github.com/eatbytes/razboynik/services/kernel"
// )

// func Delete(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
// 	var (
// 		request *razboy.REQUEST
// 		action  string
// 		scope   string
// 		err     error
// 	)

// 	scope = kc.GetArrItem(1)

// 	if scope == "" {
// 		return kc, errors.New("You should give the path of the file")
// 	}

// 	if c.Shellscope != "" {
// 		scope = c.Shellscope + "/" + scope
// 	}

// 	action = phpadapter.CreateDelete(scope) + phpadapter.CreateAnswer(c.Method, c.Parameter)
// 	request = razboy.CreateRequest(action, c)

// 	_, err = kc.Send(request)

// 	return kc, err
// }

type Deletecmd struct{}

func (delete *Deletecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		action  string
		scope   string
		args    []string
		err     error
		request *razboy.REQUEST
	)

	args = kl.GetArr()

	if len(args) < 1 {
		return delete, errors.New("You should give the path of the file to delete")
	}

	scope = args[0]

	if config.Shellscope != "" {
		scope = config.Shellscope + "/" + scope
	}

	action = "if(is_dir('" + scope + "')){$r=rmdir('" + scope + "');}else{$r=unlink('" + scope + "');}" + razboy.CreateAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)

	_, err = razboy.Send(request)

	if err != nil {
		return delete, err
	}

	kl.WriteSuccess("Delete successfully")

	return delete, nil
}

func (delete *Deletecmd) GetName() string {
	return "-delete"
}

func (delete *Deletecmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (delete *Deletecmd) GetResult() []byte {
	return make([]byte, 0)
}

func (delete *Deletecmd) GetResultStr() string {
	return ""
}
