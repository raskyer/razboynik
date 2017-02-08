package main

import (
	"os"

	"fmt"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		scope    string
		err      error
		rpc      *razboy.RPCClient
		config   *razboy.Config
		response *razboy.RESPONSE
	)

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(razboy.RPC_ERROR)
	}

	scope = getScope(os.Args[1:], config)
	response, err = list(scope, config)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(razboy.NETWORK_ERROR)
	}

	fmt.Fprintln(os.Stdout, response.GetResult())
}

func list(scope string, config *razboy.Config) (*razboy.RESPONSE, error) {
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
