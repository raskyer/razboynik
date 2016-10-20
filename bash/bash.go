package bash

import (
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eatbytes/fuzz/ferror"
	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/normalizer"
	"github.com/eatbytes/fuzz/php"
	"github.com/eatbytes/fuzz/shell"
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
	history     []*BashCommand
}

func CreateApp(srv *network.NETWORK, shl *shell.SHELL, php *php.PHP) *BashInterface {
	bsh := BashInterface{
		commonCmd: []string{"ls", "cat", "rm"},
		server:    srv,
		shell:     shl,
		php:       php,
	}

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
		b.history = append(b.history, bc)
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

func (b *BashInterface) Exit(bc *BashCommand) {
	b.Stop()
}

func (b *BashInterface) Sys(bc *BashCommand) {
	result, err := sysgo.Call(bc.str)
	bc.Write(result, err)
}

func (b *BashInterface) Encode(bc *BashCommand) {
	str := bc.str

	if str == "" {
		lastC := bc.parent.GetFormerCommand()
		str = lastC.res
	}

	sEnc := normalizer.Encode(str)
	bc.Write(sEnc, nil)
}

func (b *BashInterface) Decode(bc *BashCommand) {
	str := bc.str

	if str == "" {
		lastC := bc.parent.GetFormerCommand()
		str = lastC.res
	}

	sDec, err := normalizer.Decode(str)
	bc.Write(sDec, err)
}

func (b *BashInterface) AddSpCmd(name string, function spFunc) {
	b.specialCmd = append(b.specialCmd, name)
	b.specialFunc = append(b.specialFunc, function)
}

func (b *BashInterface) SetDefaultFunc(fn spFunc) {
	b.defaultFunc = fn
}

func (b *BashInterface) GetFormerCommand() *BashCommand {
	var lgt int
	lgt = len(b.history)

	if lgt < 2 {
		return b.history[0]
	}

	return b.history[lgt-2]
}

func (b *BashInterface) FlushHistory(bc *BashCommand) {
	b.history = nil
}

func (b *BashInterface) CreateCommand(raw string) *BashCommand {
	var arr []string
	var fnInt int
	var fn spFunc
	var strArr []string
	var str string
	var out string
	var err string

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
