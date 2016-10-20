package shellmodule

import (
	"strings"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
)

func Cd(bc *bash.BashCommand) {
	var srv *network.NETWORK
	var shl *shell.SHELL
	var result string
	var raw string
	var cd string
	var err error

	srv = bc.GetServer()
	shl = bc.GetShell()
	raw = bc.GetRaw()

	cd = shl.Cd(raw) + srv.Response()
	result, err = srv.QuickProcess(cd)

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
