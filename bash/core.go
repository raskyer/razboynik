package bash

import (
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/fuzzcore"
)

type spFunc func(*BashCommand)

type BashInterface struct {
	commonCmd   []string
	specialCmd  []string
	specialFunc []spFunc
	readline    *readline.Instance
	running     bool
}

func CreateBashApp() *BashInterface {
	bsh := BashInterface{
		commonCmd: []string{"ls", "cat", "rm"},
		specialCmd: []string{
			"-raw",
			"cd",
			"-upload",
			"-download",
			"-sys",
			"-encode",
			"-decode",
			"-info",
			"-php",
			"-exit",
		},
	}

	bsh.specialFunc = []spFunc{
		bsh.SendRawShell,
		bsh.SendCd,
		bsh.SendUpload,
		bsh.SendDownload,
		bsh.Sys,
		bsh.Encode,
		bsh.Decode,
		bsh.Info,
		bsh.SendRawPHP,
		bsh.Exit,
	}

	bsh.buildPrompt()

	return &bsh
}

func (b *BashInterface) buildPrompt() {
	autocompleter := readline.NewPrefixCompleter()
	allCmd := append(b.commonCmd, b.specialCmd...)

	for _, item := range allCmd {
		child := readline.PcItem(item)
		autocompleter.SetChildren(append(autocompleter.GetChildren(), child))
	}

	config := &readline.Config{
		Prompt:          "\033[32m•\033[0m\033[32m» [Bash]$\033[0m ",
		HistoryFile:     "/tmp/readlinebash.tmp",
		AutoComplete:    autocompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	l, err := readline.NewEx(config)

	if err != nil {
		panic(err)
	}

	b.readline = l
}

func (b *BashInterface) Start() {
	if !fuzzcore.NET.IsSetup() {
		e := fuzzcore.SetupErr()
		e.Error()
		return
	}

	b.running = true

	defer b.readline.Close()
	log.SetOutput(b.readline.Stderr())

	b.loop()
}

func (b *BashInterface) loop() {
	for b.running {
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
	cmd := b.CreateCommand(l)
	cmd.Fn(cmd)
}

func (b *BashInterface) Stop() {
	b.running = false
	fuzzcore.CMD.SetContext("")
}

func (b *BashInterface) SetPrompt(p string) {
	b.readline.SetPrompt(p)
}
