package kernel

import (
	"fmt"
	"strings"

	"github.com/eatbytes/razboy"
)

type KernelCmd struct {
	res   *razboy.RESPONSE
	scope string
	name  string
	raw   string
	arr   []string
	str   string
	out   string
	err   string
}

func CreateCmd(raw string, opt ...string) *KernelCmd {
	var (
		arr                        []string
		str, out, err, name, scope string
	)

	scope = ""
	out = "&1"
	err = "&2"

	arr = strings.Fields(raw)

	if len(arr) > 0 {
		name = arr[0]

		tmp := append(arr[1:], arr[len(arr):]...)
		str = strings.Join(tmp, " ")
	}

	if len(opt) > 0 {
		scope = opt[0]
	}

	if len(opt) > 1 {
		out = opt[1]
	}

	if len(opt) > 2 {
		err = opt[2]
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
	str = strings.Trim(str, "\n")

	if kc.out == "&1" && str != "" {
		fmt.Println(str)
	}
}

func (kc KernelCmd) WriteError(err error) {
	if kc.err == "&2" {
		fmt.Println(err.Error())
	}
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

func (kc KernelCmd) GetResponse() *razboy.RESPONSE {
	return kc.res
}

func (kc KernelCmd) GetScope() string {
	return kc.scope
}

func (kc *KernelCmd) GetResult() string {
	return kc.res.GetResult()
}

func (kc *KernelCmd) SetResponse(rzRes *razboy.RESPONSE) {
	kc.res = rzRes
}

func (kc *KernelCmd) SetScope(scope string) {
	kc.scope = scope
}

func (kc *KernelCmd) Send(request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	var (
		response *razboy.RESPONSE
		err      error
	)

	response, err = razboy.Send(request)
	kc.SetResponse(response)

	return response, err
}
