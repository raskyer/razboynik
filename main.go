package main

import (
	"os"

	"github.com/eatbytes/razboynik/app"
)

func main() {
	a := app.Create()
	a.Run(os.Args)
}
