package main

import "fuzzer/src/app"

func main() {
	main := app.CreateMainApp()
	main.Start()
}
