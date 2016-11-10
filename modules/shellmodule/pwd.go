package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
)

func Pwd(bc *bash.BashCommand) {
	var (
		result, raw, pwd, line string
		err                    error
		n                      *network.NETWORK
		s                      *shell.SHELL
	)

	n, s, _ = bc.GetObjects()
	raw = bc.GetRaw()
	pwd = s.Raw(raw) + n.Response()
	result, err = n.QuickProcess(pwd)

	if err != nil {
		bc.WriteError(err)
		return
	}

	line = strings.TrimSpace(result)
	bc.GetParent().UpdatePrompt(line)

	bc.WriteSuccess(result)
}
