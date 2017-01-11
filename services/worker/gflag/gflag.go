package gflag

import (
	"github.com/eatbytes/razboy"
)

var (
	Method      string
	Parameter   string
	Key         string
	Shellmethod string
	Shellscope  string
	Proxy       string
	Debug       bool
	Silent      bool
	Noextra     bool
	Rpc         bool
	Encoding    int
)

func BuildConfig(url string) *razboy.Config {
	return &razboy.Config{
		Method:      razboy.MethodToInt(Method),
		Parameter:   Parameter,
		Key:         Key,
		Proxy:       Proxy,
		Encoding:    Encoding,
		Shellmethod: razboy.ShellmethodToInt(Shellmethod),
		Shellscope:  Shellscope,
	}
}
