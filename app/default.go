package app

import (
	"github.com/eatbytes/fuzz/core"
	"github.com/eatbytes/fuzzer/printer"
	"github.com/urfave/cli"
)

func (main *MainInterface) Default(c *cli.Context) {
	printer.PrintSection("Configuration", "Configure the option of the remote server")

	url := read("Url")
	method := readInt("Method")
	parameter := read("Parameter")
	shmethod := read("PHP Method")

	cf := core.Config{
		url,
		method,
		parameter,
		shmethod,
		false,
	}

	main.startProcess(&cf)
}

func read(str string) string {
	var input string
	var err error

	printer.Println(str + " : ")
	printer.Print("-> ")
	input, err = printer.Read()

	if err != nil {
		printer.Error(err)
		return read(str)
	}

	return input
}

func readInt(str string) int {
	var input int
	var err error

	printer.Println(str + " : ")
	printer.Print("-> ")
	input, err = printer.ReadInt()

	if err != nil {
		printer.Error(err)
		return readInt(str)
	}

	return input
}
