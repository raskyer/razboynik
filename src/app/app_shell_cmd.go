package app

import (
	"fuzzer"
	"fuzzer/src/common"
	"strings"

	"github.com/urfave/cli"
)

func (main *MainInterface) SendLs(c *cli.Context) {
	var ls string

	if c.Bool("raw") {
		r := "ls " + strings.Join(c.Args(), " ")
		ls = fuzzer.CMD.Raw(r)
	} else {
		context := c.Args().Get(0)
		ls = fuzzer.CMD.Ls(context)
	}

	result, err := common.Process(ls)

	if err != nil {
		err.Error()
		return
	}

	common.ReadEncode(result)
}
