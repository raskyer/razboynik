package modules

import (
	"github.com/eatbytes/fuzzer/bash"
	"github.com/eatbytes/fuzzer/modules/shellmodule"
)

func Boot(b *bash.BashInterface) {
	b.AddSpCmd("-raw", shellmodule.Raw)
}
