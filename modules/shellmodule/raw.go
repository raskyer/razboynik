package shellmodule

import (
	"github.com/eatbytes/fuzz/normalizer"
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/processor"
)

func Raw(bc *bash.BashCommand) {
	raw := bc.GetShell().Raw(bc.GetRaw())
	srv := bc.GetServer()
	result, err := processor.Process(srv, raw)

	if err != nil {
		bc.WriteError(err)
		return
	}

	result, err = normalizer.Decode(result)

	bc.Write(result, err)
}
