package main

import (
	"os"

	"github.com/eatbytes/fuzzer/app"
)

func main() {
	var main *app.MainInterface

	main = app.CreateApp()
	main.Run(os.Args)
}
