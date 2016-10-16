package main

import (
	"os"

	"github.com/eatbytes/fuzzer/app"
)

func main() {
	main := app.CreateMainApp()
	main.Run(os.Args)
}
