package app

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

func (app *AppInterface) Default(c *cli.Context) {
	var url, method, parameter, key string
	var shmethod int

	services.PrintSection("Configuration", "Configure the option of the remote server")

	url = read("URL*", true)
	method = read("METHOD", false)
	parameter = read("PARAMETER", false)
	shmethod = readInt("PHP METHOD")
	key = read("KEY", false)

	cf := core.Config{
		url,
		method,
		parameter,
		shmethod,
		key,
		false,
	}

	err := app.startProcess(&cf)

	if err != nil {
		services.PrintError(err)
	}
}

func read(str string, required bool) string {
	var input string
	var err error

	services.Println(str + " : ")
	services.Print("-> ")
	input, err = services.Read()

	if err != nil {
		services.PrintError(err)
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

	services.Println(str + " : ")
	services.Print("-> ")
	input, err = services.ReadInt()

	if err != nil {
		services.PrintError(err)
		return readInt(str)
	}

	return input
}
