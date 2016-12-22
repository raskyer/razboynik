// package phpmodule

// import (
// 	"github.com/eatbytes/razboy"
// 	"github.com/eatbytes/razboynik/services/kernel"
// )

// func Raw(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
// 	var (
// 		request *razboy.REQUEST
// 		err     error
// 	)

// 	request = razboy.CreateRequest(kc.GetStr(), c)
// 	_, err = kc.Send(request)

// 	return kc, err
// }
