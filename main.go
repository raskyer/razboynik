package main

import (
	"os"

	"github.com/eatbytes/razboynik/app"
)

func main() {
	var a *app.AppInterface

	a = app.Create()
	a.Run(os.Args)
}
