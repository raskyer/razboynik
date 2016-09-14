package bash

import (
	"command"
	"core"
	"global"
	"strings"

	"github.com/chzyer/readline"
)

type spFunc func(string)

var BSH = BashInterface{
	spCmd:     []string{"cd", "vim"},
	spCmdFunc: []spFunc{command.CMD.Cd},
}

type BashInterface struct {
	spCmd     []string
	spCmdFunc []spFunc
	prompt    *string
}

func (b *BashInterface) GetPrompt() *readline.Instance {
	config := &readline.Config{
		Prompt:          "\033[31m»\033[0m [Bash]$ ",
		HistoryFile:     "/tmp/readlinebash.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	l, err := readline.NewEx(config)

	if err != nil {
		panic(err)
	}

	global.Global.BashReadline = l

	return l
}

func (b *BashInterface) Run(l string) {
	if l == "exit" {
		core.ExitBash()
		return
	}

	arr := strings.Fields(l)
	for i, item := range b.spCmd {
		if item == arr[0] {
			b.spCmdFunc[i](l)
			return
		}
	}

	command.CMD.Raw(l)
}
