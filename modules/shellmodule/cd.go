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

	srv, shl, _ = bc.GetObjects()
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

		bc.GetParent().UpdatePrompt(line)
		bc.WriteSuccess(result)
	}
}