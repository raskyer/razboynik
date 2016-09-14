package global

import "github.com/chzyer/readline"

var Global = GlobalInterface{}

type GlobalInterface struct {
	BashReadline *readline.Instance
}
