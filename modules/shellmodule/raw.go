package shellmodule

import (
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
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
