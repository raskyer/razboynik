package bash

import (
	"io"
	"log"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eatbytes/fuzz/ferror"
	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/normalizer"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/sysgo"
)

type spFunc func(*BashCommand)

type BashInterface struct {
	commonCmd   []string
	specialCmd  []string
	specialFunc []spFunc
	readline    *readline.Instance
	running     bool
	server      *network.NETWORK
	shell       *shell.SHELL
	history     []BashCommand
}

func CreateBashApp(srv *network.NETWORK, shl *shell.SHELL) *BashInterface {
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
		server: srv,
		shell:  shl,
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
	if !b.server.IsSetup() {
		e := ferror.SetupErr()
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

		bc := b.CreateCommand(line)
		b.history = append(b.history, *bc)
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

func (b *BashInterface) Exit(cmd *BashCommand) {
	b.Stop()
}

func (b *BashInterface) Sys(cmd *BashCommand) {
	result, err := sysgo.Call(cmd.str)
	cmd.Write(result, err)
}

func (b *BashInterface) Encode(cmd *BashCommand) {
	sEnc := normalizer.Encode(cmd.str)
	cmd.WriteSuccess(sEnc)
}

func (b *BashInterface) Decode(cmd *BashCommand) {
	str := cmd.str

	if cmd.str != "" {
		str = ""
	}

	sDec, err := normalizer.Decode(str)
	cmd.Write(sDec, err)
}

func (b *BashInterface) AddSpCmd(name string, function spFunc) {
	b.specialCmd = append(b.specialCmd, name)
	b.specialFunc = append(b.specialFunc, function)
}

func (b *BashInterface) CreateCommand(raw string) *BashCommand {
	arr := strings.Fields(raw)

	fnInt := defineFunc(arr[0], b.specialCmd)
	fn := b.specialFunc[fnInt]

	strArr := append(arr[1:], arr[len(arr):]...)
	str := strings.Join(strArr, " ")

	out := defineOutput(raw, arr)
	err := defineErrput(raw, arr)

	cmd := BashCommand{
		raw:    raw,
		arr:    arr,
		str:    str,
		out:    out,
		err:    err,
		fn:     fn,
		parent: b,
	}

	return &cmd
}
