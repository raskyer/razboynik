package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razpomoshnik"
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
		return
	}

	raw = "pwd " + strings.Join(os.Args[1:], " ")
	a = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod) + razboy.AddAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(a, config)
	response, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	scope = strings.TrimSpace(response.GetResult())

	if scope != "" {
		err = razpomoshnik.UpdatePrompt(scope)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
	}
}
