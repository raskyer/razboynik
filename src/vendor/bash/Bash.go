package bash

import (
	"command"
	"core"
	"strings"

	"github.com/chzyer/readline"
)

type spFunc func(string)

type BashInterface struct {
	spCmd     []string
	spCmdFunc []spFunc
}

func CreateBash() *BashInterface {
	b := BashInterface{
		spCmd:     []string{"cd", "vim"},
		spCmdFunc: []spFunc{command.CMD.Cd},
	}

	return &b
}

func (b *BashInterface) GetPrompt() *readline.Instance {
	l, err := readline.NewEx(&readline.Config{
		Prompt:          "\033[31m»\033[0m [Bash]$ ",
		HistoryFile:     "/tmp/readlinebash.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	})

	if err != nil {
		panic(err)
	}

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
