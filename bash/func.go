package bash

import (
	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/sysgo"
)

func (b *BashInterface) Exit(cmd *BashCommand) {
	b.Stop()
}

func (b *BashInterface) Sys(cmd *BashCommand) {
	result, err := sysgo.Call(cmd.str)
	cmd.Write(result, err)
}

func (b *BashInterface) Encode(cmd *BashCommand) {
	sEnc := fuzzcore.Encode(cmd.str)
	cmd.WriteSuccess(sEnc)
}

func (b *BashInterface) Decode(cmd *BashCommand) {
	str := cmd.str

	if cmd.str != "" {
		str = ""
	}

	sDec, err := fuzzcore.Decode(str)
	cmd.Write(sDec, err)
}
