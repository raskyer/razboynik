package shellmodule

import (
	"strings"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/normalizer"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/processor"
)

func Cd(bc *bash.BashCommand) {
	var srv *network.NETWORK
	var shl *shell.SHELL

	srv = bc.GetServer()
	shl = bc.GetShell()

	cd := shl.Cd(bc.GetRaw())
	result, err := processor.Process(srv, cd)

	if err != nil {
		bc.WriteError(err)
		return
	}

	result, err = normalizer.Decode(result)

	if err != nil {
		bc.WriteError(err)
		return
	}

	line := strings.TrimSpace(result)

	if line != "" {
		shl.SetContext(line)

		bc.GetParent().SetPrompt("\033[32m•\033[0m\033[32m» [Bash]:" + line + "$\033[0m ")
		bc.WriteSuccess(result)
	}
}
