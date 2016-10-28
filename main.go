package main

import (
	"os"

	"github.com/eatbytes/razboynik/app"
)

func main() {
	var appli *app.AppInterface

	appli = app.CreateApp()
	appli.Run(os.Args)
}
