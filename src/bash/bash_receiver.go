package bash

import (
	"fmt"
	"fuzzer"
	"strings"
)

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzer.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzer.CMD.SetContext(line)
		b.SetPrompt("\033[32m•\033[0m\033[32m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}
