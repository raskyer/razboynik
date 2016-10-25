package shellmodule

import (
	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
)

func Raw(bc *bash.BashCommand) {
	var shl *shell.SHELL
	var srv *network.NETWORK
	var raw string
	var result string
	var r string
	var err error

	srv, shl, _ = bc.GetObjects()

	raw = bc.GetRaw()
	r = shl.Raw(raw) + srv.Response()

	result, err = srv.QuickProcess(r)

	bc.Write(result, err)
}
