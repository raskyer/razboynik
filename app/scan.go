package app

import (
	"errors"
	"strconv"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

func (app *AppInterface) Scan(c *cli.Context) {
	var (
		url, parameter, key string
		mPool               []string
		sPool               []int
		sDef, lgt, round    int
		raw                 bool
		err                 error
		cf                  *core.Config
	)

	url = c.Args().First()
	parameter = c.String("p")
	key = c.String("k")

	if url == "" {
		err = errors.New("Flag -u (url) is required")
		services.PrintError(err)
		return
	}

	services.PrintSection("Scanning", "Scan target: "+url+", parameter: "+parameter)

	mPool = []string{"GET", "POST", "HEADER", "COOKIE"}
	sPool = []int{0, 1}
	sDef = 0
	round = 0
	raw = false
	lgt = len(mPool)

	for i := 0; i < lgt; i++ {
		cf = &core.Config{
			Url:       url,
			Method:    mPool[i],
			Parameter: parameter,
			Shmethod:  sPool[sDef],
			Key:       key,
			Raw:       raw,
			Crypt:     false,
		}

		err = app.testing(cf)
		app.printscan(cf, err)

		round++
		if round == 4 {
			services.Println("")
			round = 0
		}

		if !raw {
			raw = true
			i = i - 1
			continue
		}

		raw = false

		if sDef == 0 {
			sDef = 1
			i = i - 1
			continue
		}

		sDef = 0
	}
}

func (app *AppInterface) printscan(cf *core.Config, err error) {
	var str, shell string

	if err != nil {
		str = color.RedString("[x] ")
	} else {
		str = color.GreenString("[v] ")
	}

	if cf.Shmethod == 1 {
		shell = "shell_exec()"
	} else {
		shell = "system()"
	}

	str = str + "Method: " + color.YellowString(cf.Method) + ", Shell method: " + color.MagentaString(shell) + ", Raw: " + color.BlueString(strconv.FormatBool(cf.Raw))
	services.Println(str)
}
