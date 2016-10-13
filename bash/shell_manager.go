package bash

import (
	"fmt"
	"strings"

	"github.com/eatbytes/fuzzcore"
	"github.com/leaklessgfy/fuzzer/networking"
	"github.com/leaklessgfy/fuzzer/reader"
)

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzcore.CMD.Raw(str)
	result, err := networking.Process(raw)

	if err != nil {
		err.Error()
		return
	}

	reader.ReadEncode(result)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzcore.CMD.Cd(str)
	result, err := networking.Process(cd)

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
