package app

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services"
	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type scanobj struct {
	url    string
	key    string
	method string
	shellm string
	raw    bool
	status bool
}

func (app *AppInterface) Scan(c *cli.Context) {
	var url, parameter, key, shell string
	var mPool []string
	var sPool []int
	var sDef int
	var raw bool
	var err error
	var cf *core.Config

	url = c.String("u")
	parameter = c.String("p")
	key = c.String("k")

	mPool = []string{"GET", "POST", "HEADER", "COOKIE"}
	sPool = []int{0, 1}
	sDef = 0
	raw = false

	for i := 0; i < 3; i++ {
		cf = &core.Config{
			url,
			mPool[i],
			parameter,
			sPool[sDef],
			key,
			raw,
			false,
		}

		if sDef == 1 {
			shell = "shell_exec()"
		} else {
			shell = "system()"
		}

		sc := scanobj{
			url,
			key,
			mPool[i],
			shell,
			raw,
			true,
		}

		err = app.testing(cf)

		if err != nil {
			if !raw {
				color.Red("Method: "+mPool[i]+", Raw: False, sDef: %d", sDef)
				raw = true
				i = i - 1
				continue
			}

			color.Red("Method: "+mPool[i]+", Raw: True, sDef: %d", sDef)

			raw = false

			if sDef == 0 {
				sDef = 1
				i = i - 1
				continue
			}

			sDef = 0
		}
	}
}

func (app *AppInterface) printscan(sc scanobj) {
	var str string

	if sc.status {
		color.RedString("[x] URL: "+sc.url+", Method: "+sc.method+", Shell Method: "+sc.shellm+", Raw: %v", sc.raw)
		services.Println(str)
	}
}
