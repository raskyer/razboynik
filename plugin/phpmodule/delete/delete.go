package main

import (
	"fmt"
	"os"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		action  string
		scope   string
		err     error
		rpc     *razboy.RPCClient
		config  *razboy.Config
		request *razboy.REQUEST
	)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You should give the path of the file to delete")
		return
	}

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()
	scope = os.Args[1]

	if config.Shellscope != "" {
		scope = config.Shellscope + "/" + scope
	}

	action = "if(is_dir('" + scope + "')){$r=rmdir('" + scope + "');}else{$r=unlink('" + scope + "');}" + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)

	_, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Fprintln(os.Stdout, "Delete successfully")
}
