package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/common"
	"io"
	"log"
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
		spCmd: []string{"exit", "cd", "-upload", "-download", "-sys"},
	}

	app.spCmdFunc = []spFunc{app.Exit, app.SendCd, app.SendUpload, app.SendDownload, app.Sys}
	app._buildPrompt()

	return &app
}

func (b *BashInterface) _buildPrompt() {
	config := &readline.Config{
		Prompt:          "\033[32m•\033[0m\033[32m»\033[0m [Bash]$ ",
		HistoryFile:     "/tmp/readlinebash.tmp",
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	l, err := readline.NewEx(config)

	if err != nil {
		panic(err)
	}

	b.readline = l
}

func (b *BashInterface) loop() {
	defer b.readline.Close()
	log.SetOutput(b.readline.Stderr())

	for b.IsRunning() {
		line, err := b.readline.Readline()
		if err == readline.ErrInterrupt || err == io.EOF {
			return
		}

		if len(line) == 0 {
			continue
		}

		b.Run(line)
	}
}

func (b *BashInterface) Run(l string) {
	if strings.Contains(l, "&&") {
		b.SendRaw(l)
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

func (b *BashInterface) Start() {
	if !fuzzer.NET.IsSetup() {
		fmt.Println("You haven't setup the required information, please refer to srv config")
		return
	}

	b.running = true
	b.loop()
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

func (b *BashInterface) Exit(str string) {
	b.Stop()
}

func (b *BashInterface) Sys(str string) {
	arr := strings.Fields(str)

	if len(arr) < 2 {
		return
	}

	arr = append(arr[1:], arr[len(arr):]...)
	full := strings.Join(arr, " ")

	common.Syscall(full)
}
