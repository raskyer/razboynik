package shell

import (
	"fmt"
	"fuzzer/src/common"
	"strings"

	"fuzzer/src/bash"
	"github.com/eatbytes/fuzzcore"
)

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzcore.CMD.Raw(str)
	result, err := common.Process(raw)

	if err != nil {
		err.Error()
		return
	}

	common.ReadEncode(result)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzcore.CMD.Cd(str)
	result, err := common.Process(cd)

	if err != nil {
		err.Error()
		return
	}

	b.ReceiveCd(result)
}

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzcore.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzcore.CMD.SetContext(line)
		b.SetPrompt("\033[32m•\033[0m\033[32m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}
