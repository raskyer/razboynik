package kernelmodule

import (
	"errors"
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
	)

	fkc = kernel.Boot().GetFormerCmd()

	if fkc == nil || fkc.GetResponse() == nil || fkc.GetResponse().GetRequest() == nil {
		return kc, errors.New("You havn't made any request.\nYou must make a request before seing any information")
	}

	if !strings.Contains(fkc.GetStr(), "response") {
		kc.WriteSuccess(color.YellowString("--- Request ---"))
		b, _ = httputil.DumpRequestOut(fkc.GetResponse().GetRequest().GetHTTP(), false)
		kc.WriteSuccess(string(b))
		kc.WriteSuccess(string(fkc.GetResponse().GetRequest().GetBody()) + "\n")
	}

	if !strings.Contains(fkc.GetStr(), "request") {
		kc.WriteSuccess(color.YellowString("--- Response ---"))
		b, _ = httputil.DumpResponse(fkc.GetResponse().GetHTTP(), false)
		kc.WriteSuccess(string(b))
	}

	return fkc, nil
}
