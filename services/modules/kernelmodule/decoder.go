package kernelmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Decodeitem = kernel.Item{
	Name: "-decode",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		str := strings.Join(l.GetArg(), " ")
		sDec, err := razboy.Decode(str)

		kernel.Write(l.GetStdout(), l.GetStderr(), err, sDec)

		return kernel.Response{Err: err, Body: sDec}
	},
}

var Encodeitem = kernel.Item{
	Name: "-encode",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		str := strings.Join(l.GetArg(), " ")
		sEnc := razboy.Encode(str)

		kernel.WriteSuccess(l.GetStdout(), sEnc)

		return kernel.Response{Body: sEnc}
	},
}
