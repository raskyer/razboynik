package kernelmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Decodecmd struct{}
type Encodecmd struct{}

func (d *Decodecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	str := strings.Join(kl.GetArr(), " ")
	sDec, err := razboy.Decode(str)

	kl.Write(err, sDec)

	return d, nil
}

func (d *Decodecmd) GetName() string {
	return "decode"
}

func (d *Decodecmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (d *Decodecmd) GetResult() []byte {
	return make([]byte, 0)
}

func (d *Decodecmd) GetResultStr() string {
	return ""
}

func (e *Encodecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	str := strings.Join(kl.GetArr(), " ")
	sEnc := razboy.Encode(str)

	kl.WriteSuccess(sEnc)

	return e, nil
}

func (e *Encodecmd) GetName() string {
	return "encode"
}

func (e *Encodecmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (e *Encodecmd) GetResult() []byte {
	return make([]byte, 0)
}

func (e *Encodecmd) GetResultStr() string {
	return ""
}
