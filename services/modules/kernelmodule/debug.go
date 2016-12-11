package kernelmodule

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"net/http/httputil"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/fatih/color"
)

func Debug(kc *kernel.KernelCmd, config *razboy.Config) (*kernel.KernelCmd, error) {
	var fkc *kernel.KernelCmd

	fkc = kernel.Boot().GetFormerCmd()

	if fkc == nil || fkc.GetResponse() == nil || fkc.GetResponse().GetRequest() == nil {
		return kc, errors.New("You havn't made any request.\nYou must make a request before seing any information")
	}

	if !strings.Contains(fkc.GetStr(), "response") {
		b, _ := httputil.DumpRequestOut(fkc.GetResponse().GetRequest().GetHTTP(), true)
		fmt.Println(string(b))
		//_requestInfo(kc, fkc)
	}

	if !strings.Contains(fkc.GetStr(), "request") {
		b, _ := httputil.DumpResponse(fkc.GetResponse().GetHTTP(), false)
		fmt.Println(string(b))
		//_responseInfo(kc, fkc)
	}

	return fkc, nil
}

func _requestInfo(kc, fkc *kernel.KernelCmd) {
	var (
		flag bool
		str  string
		r    *http.Request
	)

	color.Yellow("--- Request ---")

	flag = false
	r = fkc.GetResponse().GetRequest().GetHTTP()
	str = kc.GetStr()

	if strings.Contains(str, "-url") {
		color.Cyan("Url: ")
		fmt.Println(r.URL.String())
		flag = true
	}

	if strings.Contains(str, "-method") {
		color.Cyan("Method: ")
		fmt.Println(r.Method)
		flag = true
	}

	if strings.Contains(str, "-body") {
		color.Cyan("Body: ")
		fmt.Println(r.PostForm)
		flag = true
	}

	if strings.Contains(str, "-header") {
		color.Cyan("Header: ")
		fmt.Println(r.Header)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}

	color.Unset()
}

func _responseInfo(kc, fkc *kernel.KernelCmd) {
	var (
		flag bool
		str  string
		r    *http.Response
	)

	color.Yellow("--- Response ---")

	flag = false
	r = fkc.GetResponse().GetHTTP()
	str = kc.GetStr()

	if strings.Contains(str, "-status") {
		color.Cyan("Status:")
		fmt.Println(r.Status)
		flag = true
	}

	if strings.Contains(str, "-body") {
		color.Cyan("Body:")
		fmt.Println(fkc.GetResult())
		flag = true
	}

	if strings.Contains(str, "-headers") {
		color.Cyan("Headers:")
		fmt.Println(r.Header)
		flag = true
	}

	if !flag {
		fmt.Println(r)
	}

	color.Unset()
}
