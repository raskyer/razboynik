package phpmodule

import "github.com/eatbytes/fuzzer/bash/networking"

func Raw(bc *BashCommand) {
	raw := bc.parent.PHP.Raw(cmd.str)
	result, err := networking.Process(raw)
	cmd.Write(result, err)
}
