package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/reader"
	"strings"
)

func (b *BashInterface) Exit(str string) {
	b.Stop()
}

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzer.CMD.Raw(str)
	fuzzer.NET.Send(raw, reader.ReadEncode)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	fuzzer.NET.Send(cd, b.ReceiveCd)
}

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzer.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzer.CMD.SetContext(line)
		b.SetPrompt("\033[31m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}
