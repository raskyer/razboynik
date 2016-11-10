package bash

import (
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy/ferror"
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/normalizer"
	"github.com/eatbytes/razboy/php"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/sysgo"
)

type spFunc func(*BashCommand)

type BashInterface struct {
	commonCmd   []string
	specialCmd  []string
	specialFunc []spFunc
	defaultFunc spFunc
	readline    *readline.Instance
	running     bool
	server      *network.NETWORK
	shell       *shell.SHELL
	php         *php.PHP
}

func Create(srv *network.NETWORK, shl *shell.SHELL, php *php.PHP) *BashInterface {
	return &BashInterface{
		commonCmd: []string{"ls", "cat", "rm"},
		server:    srv,
		shell:     shl,
		php:       php,
	}
}

func (b *BashInterface) buildPrompt() {
	autocompleter := readline.NewPrefixCompleter()
	allCmd := append(b.commonCmd, b.specialCmd...)

	for _, item := range allCmd {
		child := readline.PcItem(item)
		autocompleter.SetChildren(append(autocompleter.GetChildren(), child))
	}

	config := &readline.Config{
		Prompt:          "(" + b.server.GetUrl() + ")$ ",
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
	if !b.server.IsSetup() {
		e := ferror.SetupErr()
		e.Error()
		return
	}

	b.buildPrompt()

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

		bc := b.CreateCommand(line)
		bc.Exec()
	}
}

func (b *BashInterface) Run(bc *BashCommand) {
	bc.Exec()
}

func (b *BashInterface) Stop() {
	b.running = false
	b.shell.SetContext("")
}

func (b *BashInterface) SetPrompt(p string) {
	b.readline.SetPrompt(p)
}

func (b *BashInterface) UpdatePrompt(p string) {
	b.readline.SetPrompt("(" + b.server.GetUrl() + "):" + p + "$ ")
}

func (b *BashInterface) Exit(bc *BashCommand) {
	b.Stop()
}

func (b *BashInterface) Sys(bc *BashCommand) {
	result, err := sysgo.Call(bc.str)
	bc.Write(result, err)
}

func (b *BashInterface) Encode(bc *BashCommand) {
	sEnc := normalizer.Encode(bc.str)
	bc.Write(sEnc, nil)
}

func (b *BashInterface) Decode(bc *BashCommand) {
	if bc.str == "" {
		return
	}

	sDec, err := normalizer.Decode(bc.str)
	bc.Write(sDec, err)
}

func (b *BashInterface) AddSpCmd(name string, function spFunc) {
	b.specialCmd = append(b.specialCmd, name)
	b.specialFunc = append(b.specialFunc, function)
}

func (b *BashInterface) SetDefaultFunc(fn spFunc) {
	b.defaultFunc = fn
}

func (b *BashInterface) CreateCommand(raw string) *BashCommand {
	var (
		arr    []string
		fnInt  int
		fn     spFunc
		strArr []string
		str    string
		out    string
		err    string
	)

	arr = strings.Fields(raw)

	fnInt = defineFunc(arr[0], b.specialCmd)

	if fnInt == -1 {
		fn = b.defaultFunc
	} else {
		fn = b.specialFunc[fnInt]
	}

	strArr = append(arr[1:], arr[len(arr):]...)
	str = strings.Join(strArr, " ")

	out = defineOutput(raw, arr)
	err = defineErrput(raw, arr)

	return &BashCommand{
		raw:    raw,
		arr:    arr,
		str:    str,
		out:    out,
		err:    err,
		fn:     fn,
		parent: b,
	}
}
