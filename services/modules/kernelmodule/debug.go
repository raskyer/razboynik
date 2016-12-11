package kernelmodule

import (
	"errors"
	"fmt"
	"strings"

	"net/http/httputil"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/fatih/color"
)

func Debug(kc *kernel.KernelCmd, config *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		fkc *kernel.KernelCmd
		b   []byte
		err error
	)

	fkc = kernel.Boot().GetFormerCmd()

	if fkc == nil || fkc.GetResponse() == nil || fkc.GetResponse().GetRequest() == nil {
		return kc, errors.New("You havn't made any request.\nYou must make a request before seing any information")
	}

	if !strings.Contains(fkc.GetStr(), "response") {
		color.Yellow("--- Request ---")
		b, err = httputil.DumpRequestOut(fkc.GetResponse().GetRequest().GetHTTP(), true)

		if err != nil {
			b, _ = httputil.DumpRequestOut(fkc.GetResponse().GetRequest().GetHTTP(), false)
			color.Red(err.Error())
		}

		fmt.Println(string(b))
	}

	if !strings.Contains(fkc.GetStr(), "request") {
		color.Yellow("--- Response ---")
		b, err = httputil.DumpResponse(fkc.GetResponse().GetHTTP(), false)

		if err != nil {
			color.Red(err.Error())
		}

		fmt.Println(string(b))
	}

	return fkc, nil
}
