// package phpmodule

// import (
// 	"github.com/eatbytes/razboy"
// 	"github.com/eatbytes/razboy/adapter/phpadapter"
// 	"github.com/eatbytes/razboynik/services/kernel"
// )

// func ListFile(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
// 	var (
// 		request *razboy.REQUEST
// 		action  string
// 		scope   string
// 		err     error
// 	)

// 	scope = "__DIR__"

// 	if c.Shellscope != "" {
// 		scope = "'" + c.Shellscope + "'"
// 	}

// 	if kc.GetStr() != "" {
// 		scope = "'" + kc.GetStr() + "'"
// 	}

// 	action = phpadapter.CreateListFile(scope) + phpadapter.CreateAnswer(c.Method, c.Parameter)
// 	request = razboy.CreateRequest(action, c)

// 	_, err = kc.Send(request)

// 	return kc, err
// }
