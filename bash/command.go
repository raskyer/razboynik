package bash

import (
	"fmt"
	"strings"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/shell"
)

type BashCommand struct {
	raw    string
	arr    []string
	str    string
	out    string
	err    string
	fn     spFunc
	res    string
	code   int
	parent *BashInterface
}

func (bc *BashCommand) Write(str string, err error) {
	if err != nil {
		bc.WriteError(err)
		return
	}

	bc.WriteSuccess(str)
}

func (bc *BashCommand) WriteSuccess(str string) {
	if bc.out == "1" {
		fmt.Println(str)
	}
}

func (bc *BashCommand) WriteError(err error) {
	if bc.err == "2" {
		fmt.Println(err.Error())
	}
}

func (bc *BashCommand) Exec() {
	bc.fn(bc)
}

func (bc *BashCommand) GetParent() *BashInterface {
	return bc.parent
}

func (bc *BashCommand) GetServer() *network.NETWORK {
	return bc.parent.server
}

func (bc *BashCommand) GetShell() *shell.SHELL {
	return bc.parent.shell
}

func (bc *BashCommand) GetRaw() string {
	return bc.raw
}

func (bc *BashCommand) GetStr() string {
	return bc.str
}

func defineOutput(str string, arr []string) string {
	if strings.Contains(str, ">") {

	}

	return "1"
}

func defineErrput(str string, arr []string) string {
	if strings.Contains(str, "2>") {

	}

	return "2"
}

func defineFunc(str string, cmds []string) int {
	for i, item := range cmds {
		if str == item {
			return i
		}
	}

	return 0
}
