package shellmodule

import (
	"strings"

	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
)

func Cd(bc *bash.BashCommand) {
	var (
		result, raw, cd, line string
		err                   error
		n                     *network.NETWORK
		s                     *shell.SHELL
	)

	n, s, _ = bc.GetObjects()
	raw = bc.GetRaw()

	if strings.Contains(raw, "&&") || bc.GetArrItem(1, "") == "-" {
		Raw(bc)
		return
	}

	cd = s.Cd(raw) + n.Response()
	result, err = n.QuickProcess(cd)

	if err != nil {
		bc.WriteError(err)
		return
	}

	line = strings.TrimSpace(result)

	if line != "" {
		s.SetContext(line)
		bc.GetParent().UpdatePrompt(line)

		bc.WriteSuccess(result)
	}
}
