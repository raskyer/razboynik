package app

import (
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

func (app *AppInterface) Generate(c *cli.Context) {
	services.PrintSection("Generating", "Generate php file")
}
