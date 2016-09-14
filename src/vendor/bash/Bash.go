package bash

import (
	"command"
	"core"

	"github.com/chzyer/readline"
)

type BashInterface struct{}

func CreateBash() *BashInterface {
	b := BashInterface{}

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

	command.CMD.Raw(l)
}
