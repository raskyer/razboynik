package main

import "github.com/chzyer/readline"

var Global = GlobalInterface{}

type GlobalInterface struct {
	BashSession  bool
	MainSession  bool
	BashReadline *readline.Instance
}
