package shellmodule

import (
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
)

func Raw(bc *bash.BashCommand) {
	var (
		result, raw, r string
		err            error
		n              *network.NETWORK
		s              *shell.SHELL
	)

	n, s, _ = bc.GetObjects()
	raw = bc.GetRaw()
	r = s.Raw(raw) + n.Response()
	result, err = n.QuickProcess(r)

	bc.Write(result, err)
}
