package modules

import (
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/modules/phpmodule"
	"github.com/eatbytes/fuzzer/modules/shellmodule"
)

func Boot(b *bash.BashInterface) {
	b.SetDefaultFunc(shellmodule.Raw)
	b.AddSpCmd("cd", shellmodule.Cd)
	b.AddSpCmd("pwd", shellmodule.Pwd)
	b.AddSpCmd("-raw", shellmodule.Raw)
	b.AddSpCmd("-php", phpmodule.Raw)
	b.AddSpCmd("-info", Info)
	b.AddSpCmd("-sys", b.Sys)
	b.AddSpCmd("-encode", b.Encode)
	b.AddSpCmd("-decode", b.Decode)
	b.AddSpCmd("-upload", phpmodule.UploadInit)
	b.AddSpCmd("-download", phpmodule.DownloadInit)
	b.AddSpCmd("-flush-history", b.FlushHistory)
	b.AddSpCmd("-exit", b.Exit)
}
