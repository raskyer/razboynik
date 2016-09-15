package main

import (
	"fmt"
	"fuzzer"

	"github.com/chzyer/readline"
)

var Global = GlobalInterface{}

type GlobalInterface struct {
	Main         *MainInterface
	Bash         *BashInterface
	BashSession  bool
	MainSession  bool
	BashReadline *readline.Instance
}

func (g *GlobalInterface) Stop() {
	g.MainSession = false
	g.BashSession = false
	g.Main.Stop()
	g.Bash.Stop()
}

func (g *GlobalInterface) StartMain() {
	g.MainSession = true
	g.Main.Start()
	g.StopBash()
}

func (g *GlobalInterface) StartBash() {
	g.BashSession = true
	g.Bash.Start()
	g.StopMain()
}

func (g *GlobalInterface) StopMain() {
	g.MainSession = true
	g.Main.Stop()
}

func (g *GlobalInterface) StopBash() {
	g.BashSession = false
	g.Bash.Stop()
}

func (g *GlobalInterface) ReadEncode(str string) {
	sDec := fuzzer.Decode(str)
	fmt.Println(sDec)
}

func (g *GlobalInterface) Read(str string) {
	fmt.Println(str)
}
