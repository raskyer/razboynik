package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/common"
	"strings"
)

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzer.CMD.Raw(str)
	result, err := common.Process(raw)

	if err != nil {
		err.Error()
		return
	}

	common.ReadEncode(result)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	result, err := common.Process(cd)

	if err != nil {
		err.Error()
		return
	}

	b.ReceiveCd(result)
}

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzer.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzer.CMD.SetContext(line)
		b.SetPrompt("\033[32m•\033[0m\033[32m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}
