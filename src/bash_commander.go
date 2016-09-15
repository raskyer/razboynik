package main

import (
	"fmt"
	"fuzzer"
)

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzer.CMD.Raw(str)
	fuzzer.NET.Send(raw, b.Receive)
}

func (b *BashInterface) Receive(result string) {
	fmt.Println(result)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	fuzzer.NET.Send(cd, b.Receive)
}
