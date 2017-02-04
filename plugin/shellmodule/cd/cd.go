package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razpomoshnik"
)

const (
	RPC_ERROR     int = 1
	NETWORK_ERROR int = 2
)

func main() {
	var (
		raw, a, scope string
		err           error
		config        *razboy.Config
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	config, err = razpomoshnik.GetConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(RPC_ERROR)
	}

	raw = "cd " + strings.Join(os.Args[1:], " ") + " && pwd"

	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)
	request = razboy.CreateRequest(a, config)

	response, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(NETWORK_ERROR)
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope == "" {
		return
	}

	config.Shellscope = scope
	err = razpomoshnik.UpdatePrompt(scope)
	err = razpomoshnik.UpdateConfig(config)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(RPC_ERROR)
	}
}
