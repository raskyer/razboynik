package bash

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/fuzzer/bash/cleaner"
	"github.com/eatbytes/fuzzer/bash/networking"
	"github.com/eatbytes/fuzzer/bash/parser"
	"github.com/eatbytes/fuzzer/bash/syscall"
)

type spFunc func(string)

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
			"exit",
			"cd",
			"-upload",
			"-download",
			"-sys",
			"-encode",
			"-decode",
			"-info",
			"-php",
			"-keep",
		},
	}

	bsh.specialFunc = []spFunc{
		bsh.Exit,
		bsh.SendCd,
		bsh.SendUpload,
		bsh.SendDownload,
		bsh.Sys,
		bsh.Encode,
		bsh.Decode,
		bsh.Info,
		bsh.SendRawPHP,
		bsh.Keep,
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
	if strings.Contains(l, "&&") && strings.Contains(l, "cd") {
		b.SendRaw(l)
		return
	}

	arr := strings.Fields(l)
	for i, item := range b.specialCmd {
		if item == arr[0] {
			b.specialFunc[i](l)
			return
		}
	}

	b.SendRaw(l)
}

func (b *BashInterface) Start() {
	if !fuzzcore.NET.IsSetup() {
		e := fuzzcore.SetupErr()
		e.Error()
		return
	}

	b.running = true
	b.loop()
}

func (b *BashInterface) Stop() {
	b.running = false
	fuzzcore.CMD.SetContext("")
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
	full, err := parser.ParseStr(str)

	if err != nil {
		return
	}

	syscall.Syscall(full)
}

func (b *BashInterface) Encode(str string) {
	str, err := parser.ParseStr(str)

	if err != nil {
		return
	}

	sEnc := fuzzcore.Encode(str)
	fmt.Println(sEnc)
}

func (b *BashInterface) Decode(str string) {
	str, err := parser.ParseStr(str)

	if err != nil {
		return
	}

	sDec, err := fuzzcore.Decode(str)

	if err != nil {
		err.Error()
		return
	}

	fmt.Println(sDec)
}

func (b *BashInterface) Keep(str string) {
	str, err := parser.ParseStr(str)

	if err != nil {
		return
	}

	raw := fuzzcore.CMD.Raw(str)
	result, err := networking.Process(raw)

	if err != nil {
		err.Error()
		return
	}

	result, err = fuzzcore.Decode(result)

	if err != nil {
		err.Error()
		return
	}

	cleaner.SetKeeper(result)
	cleaner.Clear()
}
