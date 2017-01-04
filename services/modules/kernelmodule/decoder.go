package kernelmodule

import (
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Decodecmd struct{}
type Encodecmd struct{}

func (d *Decodecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	str := strings.Join(kl.GetArr(), " ")
	sDec, err := razboy.Decode(str)

	kernel.Write(kl.GetStdout(), kl.GetStderr(), err, sDec)

	return nil
}

func (d *Decodecmd) GetName() string {
	return "-decode"
}

func (d *Decodecmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}

func (e *Encodecmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	str := strings.Join(kl.GetArr(), " ")
	sEnc := razboy.Encode(str)

	kernel.WriteSuccess(kl.GetStdout(), sEnc)

	return nil
}

func (e *Encodecmd) GetName() string {
	return "-encode"
}

func (e *Encodecmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
