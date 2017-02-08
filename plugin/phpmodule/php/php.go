package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		action   string
		err      error
		rpc      *razboy.RPCClient
		config   *razboy.Config
		request  *razboy.REQUEST
		response *razboy.RESPONSE
	)

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(razboy.RPCERROR)
	}

	action = strings.Join(os.Args[1:], " ")
	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(razboy.NETWORKERROR)
	}

	fmt.Fprintln(os.Stdout, response.GetResult())
}
