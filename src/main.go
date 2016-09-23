package main

import (
	"fuzzer/src/app"
	"os"
)

func main() {
	main := app.CreateMainApp()
	main.Run(os.Args)
}
