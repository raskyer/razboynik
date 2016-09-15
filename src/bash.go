package main

import (
	"fmt"
	"fuzzer"
	"strings"

	"github.com/chzyer/readline"
)

type spFunc func(string)

type BashInterface struct {
	spCmd     []string
	spCmdFunc []spFunc
	readline  *readline.Instance
	running   bool
}

func CreateBashApp() *BashInterface {
	app := BashInterface{
		spCmd: []string{"exit", "cd", "vim"},
	}

	app.spCmdFunc = []spFunc{app.Exit, app.SendCd}

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

	b.readline = l

	return l
}

func (b *BashInterface) Run(l string) {
	arr := strings.Fields(l)
	for i, item := range b.spCmd {
		if item == arr[0] {
			b.spCmdFunc[i](l)
			return
		}
	}

	b.SendRaw(l)
}

func (b *BashInterface) Start() {
	if !fuzzer.NET.IsSetup() {
		fmt.Println("You haven't setup the required information, please refer to srv config")
		return
	}

	b.running = true
}

func (b *BashInterface) Stop() {
	b.running = false
	fuzzer.CMD.SetContext("")
}

func (b *BashInterface) IsRunning() bool {
	return b.running
}

func (b *BashInterface) SetPrompt(p string) {
	b.readline.SetPrompt(p)
}
