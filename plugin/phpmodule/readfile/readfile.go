package main

import (
	"fmt"
	"os"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		action   string
		file     string
		err      error
		rpc      *razboy.RPCClient
		config   *razboy.Config
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "You should give the path of the file to read")
		return
	}

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(razboy.RPC_ERROR)
	}

	file = os.Args[1]

	action = "$r=file_get_contents('" + file + "');" + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(razboy.NETWORK_ERROR)
	}

	fmt.Fprintln(os.Stdout, response.GetResult())
}
