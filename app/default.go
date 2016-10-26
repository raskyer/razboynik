package app

import (
	"github.com/eatbytes/fuzz/core"
	"github.com/eatbytes/fuzzer/printer"
	"github.com/urfave/cli"
)

func (main *MainInterface) Default(c *cli.Context) {
	printer.PrintSection("Configuration", "Configure the option of the remote server")

	url := read("Url*", true)
	method := read("Method", false)
	parameter := read("Parameter", false)
	shmethod := readInt("PHP Method")

	cf := core.Config{
		url,
		method,
		parameter,
		shmethod,
		false,
	}

	err := main.startProcess(&cf)

	if err != nil {
		printer.Error(err)
	}
}

func read(str string, required bool) string {
	var input string
	var err error

	printer.Println(str + " : ")
	printer.Print("-> ")
	input, err = printer.Read()

	if err != nil {
		printer.Error(err)
		return read(str, required)
	}

	if required && input == "\n" {
		return read(str, required)
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
