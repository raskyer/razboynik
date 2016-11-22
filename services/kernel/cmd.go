package kernel

import (
	"fmt"
	"strings"

	"github.com/eatbytes/razboy"
)

type KernelCmd struct {
	result *razboy.RazResponse
	_body  string
	scope  string
	name   string
	raw    string
	arr    []string
	str    string
	out    string
	err    string
}

func CreateCmd(raw, scope string, opt ...string) *KernelCmd {
	var (
		arr                 []string
		str, out, err, name string
	)

	out = "1"
	err = "2"

	arr = strings.Fields(raw)

	if len(arr) > 0 {
		name = arr[0]

		tmp := append(arr[1:], arr[len(arr):]...)
		str = strings.Join(tmp, " ")
	}

	if len(opt) > 0 {
		out = opt[0]
	}

	if len(opt) > 1 {
		err = opt[1]
	}

	return &KernelCmd{
		scope: scope,
		name:  name,
		raw:   raw,
		arr:   arr,
		str:   str,
		out:   out,
		err:   err,
	}
}

func (kc KernelCmd) Write(str string, err error) {
	if err != nil {
		kc.WriteError(err)
		return
	}

	kc.WriteSuccess(str)
}

func (kc KernelCmd) WriteSuccess(str string) {
	if kc.out == "1" {
		fmt.Printf(str)
	}
}

func (kc KernelCmd) WriteError(err error) {
	if kc.err == "2" {
		fmt.Println(err.Error())
	}
}

func (kc KernelCmd) GetScope() string {
	return kc.scope
}

func (kc KernelCmd) GetName() string {
	return kc.name
}

func (kc KernelCmd) GetRaw() string {
	return kc.raw
}

func (kc KernelCmd) GetStr() string {
	return kc.str
}

func (kc KernelCmd) GetArr() []string {
	return kc.arr
}

func (kc KernelCmd) GetArrLgt() int {
	return len(kc.arr)
}

func (kc KernelCmd) GetArrItem(i int, def ...string) string {
	var item string

	item = ""
	if len(kc.arr) > i {
		return kc.arr[i]
	}

	if len(def) > 0 {
		item = def[0]
	}

	return item
}

func (kc KernelCmd) GetRzResp() *razboy.RazResponse {
	return kc.result
}

func (kc *KernelCmd) GetResult() string {
	if kc._body != "" {
		return kc._body
	}

	kc._body = kc.result.GetResult()

	return kc._body
}

func (kc *KernelCmd) SetScope(scope string) {
	kc.scope = scope
}

func (kc *KernelCmd) SetResult(rzRes *razboy.RazResponse) {
	kc.result = rzRes
}
