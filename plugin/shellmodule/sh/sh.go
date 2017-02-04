package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razpomoshnik"
)

func main() {
	var (
		action, raw string
		err         error
		config      *razboy.Config
		request     *razboy.REQUEST
		response    *razboy.RESPONSE
	)

	config, err = razpomoshnik.GetConfig()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	raw = strings.Join(os.Args, " ")

	action = razboy.CreateCMD(raw, config.Shellscope, config.Shellmethod)
	action = action + razboy.AddAnswer(config.Method, config.Parameter)

	request = razboy.CreateRequest(action, config)
	response, err = razboy.Send(request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// if sh.Debug {
	// 	kernel.WriteSuccess(kl.GetStdout(), "- REQUEST")
	// 	b, _ := httputil.DumpRequestOut(request.GetHTTP(), true)
	// 	kernel.WriteSuccess(kl.GetStdout(), string(b))

	// 	kernel.WriteSuccess(kl.GetStdout(), "\n")
	// 	kernel.WriteSuccess(kl.GetStdout(), "- RESPONSE\n\n")
	// 	b, _ = httputil.DumpResponse(response.GetHTTP(), true)
	// 	kernel.WriteSuccess(kl.GetStdout(), string(b))
	// }

	fmt.Fprintln(os.Stdout, response.GetResult())
}
