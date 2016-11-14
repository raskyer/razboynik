package modules

import (
	"github.com/eatbytes/razboynik/bash"
	"github.com/eatbytes/razboynik/modules/bashmodule"
	"github.com/eatbytes/razboynik/modules/phpmodule"
	"github.com/eatbytes/razboynik/modules/shellmodule"
)

func Boot(b *bash.BashInterface) {
	b.SetDefaultFunc(shellmodule.Raw)
	b.AddSpCmd("cd", shellmodule.Cd)
	b.AddSpCmd("pwd", shellmodule.Pwd)
	b.AddSpCmd("vim", shellmodule.Vim)
	b.AddSpCmd("-raw", shellmodule.Raw)
	b.AddSpCmd("-php", phpmodule.Raw)
	b.AddSpCmd("-info", bashmodule.Info)
	b.AddSpCmd("-sys", b.Sys)
	b.AddSpCmd("-encode", b.Encode)
	b.AddSpCmd("-decode", b.Decode)
	b.AddSpCmd("-upload", phpmodule.UploadInit)
	b.AddSpCmd("-download", phpmodule.DownloadInit)
	b.AddSpCmd("exit", b.Exit)
}
