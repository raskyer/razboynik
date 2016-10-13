package main

import (
	"os"

	"github.com/leaklessgfy/fuzzer/app"
)

func main() {
	main := app.CreateMainApp()
	main.Run(os.Args)
}
