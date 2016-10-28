package phpmodule

import (
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboynik/bash"
)

func Raw(bc *bash.BashCommand) {
	var srv *network.NETWORK
	var str string
	var result string
	var err error

	srv = bc.GetServer()
	str = bc.GetStr()

	result, err = srv.QuickSend(str)
	bc.Write(result, err)
}
