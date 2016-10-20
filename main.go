package main

import (
	"os"

	"github.com/eatbytes/fuzzer/app"
)

func main() {
	main := app.CreateApp()
	main.Run(os.Args)
}
