package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		raw, a, scope string
		err           error
		rpc           *razboy.RPCClient
		config        *razboy.Config
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(razboy.RPCERROR)
	}

	raw = "pwd " + strings.Join(os.Args[1:], " ")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(razboy.NETWORKERROR)
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope == "" {
		return
	}

	err = rpc.SetPrompt(config.Url, scope)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(razboy.RPCERROR)
	}
}
