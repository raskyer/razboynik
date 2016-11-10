package phpmodule

import (
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboynik/bash"
)

func Raw(bc *bash.BashCommand) {
	var (
		str, result string
		err         error
		n           *network.NETWORK
	)

	n = bc.GetServer()
	str = bc.GetStr()
	result, err = n.QuickSend(str)

	bc.Write(result, err)
}
