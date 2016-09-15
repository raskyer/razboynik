package main

import (
	"fmt"
	"fuzzer"
	"strings"
)

func (b *BashInterface) Exit(str string) {
	Global.StartMain()
}

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzer.CMD.Raw(str)
	fuzzer.NET.Send(raw, Global.ReadEncode)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	fmt.Println(cd)
	fuzzer.NET.Send(cd, b.ReceiveCd)
}

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzer.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzer.CMD.SetContext(line)
		Global.Bash.SetPrompt("\033[31m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}
