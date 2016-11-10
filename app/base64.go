package app

import (
	"github.com/eatbytes/razboy/normalizer"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

func (app *AppInterface) Encode(c *cli.Context) {
	sEnc := normalizer.Encode(c.Args().First())
	services.Println(sEnc)
}

func (app *AppInterface) Decode(c *cli.Context) {
	if c.NArg() < 1 {
		return
	}

	sDec, err := normalizer.Decode(c.Args().First())

	if err != nil {
		services.PrintError(err)
	}

	services.Println(sDec)
}
