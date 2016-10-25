package shellmodule

import (
	"strings"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
)

func Pwd(bc *bash.BashCommand) {
	var srv *network.NETWORK
	var shl *shell.SHELL
	var result string
	var raw string
	var pwd string
	var err error

	srv, shl, _ = bc.GetObjects()
	raw = bc.GetRaw()

	pwd = shl.Raw(raw) + srv.Response()
	result, err = srv.QuickProcess(pwd)

	if err != nil {
		bc.WriteError(err)
		return
	}

	line := strings.TrimSpace(result)

	bc.GetParent().UpdatePrompt(line)
	bc.WriteSuccess(result)
}
