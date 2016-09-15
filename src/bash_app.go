package main

import (
	"strings"

	"github.com/chzyer/readline"
)

type spFunc func(string)

type BashInterface struct {
	spCmd     []string
	spCmdFunc []spFunc
	prompt    *string
	running   bool
}

func CreateBashApp() *BashInterface {
	app := BashInterface{
		spCmd: []string{"cd", "vim"},
	}

	app.spCmdFunc = []spFunc{app.SendCd}

	return &app
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

	return l
}

func (b *BashInterface) Run(l string) {
	if l == "exit" {
		b.Stop()
		return
	}

	arr := strings.Fields(l)
	for i, item := range b.spCmd {
		if item == arr[0] {
			b.spCmdFunc[i](l)
			return
		}
	}

	b.SendRaw(l)
}
