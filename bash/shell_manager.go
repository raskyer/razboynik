package bash

import (
	"strings"

	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/fuzzer/bash/networking"
)

func (b *BashInterface) SendRawShell(cmd *BashCommand) {
	raw := fuzzcore.CMD.Raw(cmd.raw)
	result, err := networking.Process(raw)

	if err != nil {
		cmd.WriteError(err)
		return
	}

	result, err = fuzzcore.Decode(result)

	cmd.Write(result, err)
}

func (b *BashInterface) SendCd(cmd *BashCommand) {
	if strings.Contains(cmd.raw, "&&") {
		b.SendRawShell(cmd)
		return
	}

	cd := fuzzcore.CMD.Cd(cmd.raw)
	result, err := networking.Process(cd)

	if err != nil {
		cmd.WriteError(err)
		return
	}

	result, err = fuzzcore.Decode(result)

	if err != nil {
		cmd.WriteError(err)
		return
	}

	line := strings.TrimSpace(result)
	if line != "" {
		fuzzcore.CMD.SetContext(line)
		b.SetPrompt("\033[32m•\033[0m\033[32m» [Bash]:" + line + "$\033[0m ")
		cmd.WriteSuccess(result)
	}
}
