package app

import (
	"errors"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/php"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
	"github.com/eatbytes/razboynik/modules"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

func (app *AppInterface) Start(c *cli.Context) {
	var (
		err error
		cf  *core.Config
	)

	cf, err = app.getConfig(c)

	if err != nil {
		services.PrintError(err)
		return
	}

	services.PrintStart()

	err = app.testing(cf)

	if err != nil {
		services.PrintError(err)
		return
	}

	services.PrintSection("Reverse shell", "Reverse shell ready!")

	app.startBash(cf)
}

func (app *AppInterface) getConfig(c *cli.Context) (*core.Config, error) {
	var (
		url, parameter, method, key string
		shmethod                    int
		raw                         bool
		err                         error
		cf                          *core.Config
	)

	url = c.String("u")
	method = c.String("m")
	parameter = c.String("p")
	shmethod = c.Int("s")
	key = c.String("k")
	raw = c.Bool("r")

	if url == "" {
		err = errors.New("Flag -u (url) is required")
		return nil, err
	}

	cf = &core.Config{
		Url:       url,
		Method:    method,
		Parameter: parameter,
		Shmethod:  shmethod,
		Key:       key,
		Raw:       raw,
		Crypt:     false,
	}

	return cf, nil
}

func (app *AppInterface) testing(cf *core.Config) error {
	var (
		status bool
		err    error
		n      *network.NETWORK
	)

	n, err = network.Create(cf)

	if err != nil {
		return err
	}

	status, err = n.Test()

	if err != nil || status != true {
		return err
	}

	return nil
}

func (app *AppInterface) startBash(cf *core.Config) {
	var (
		n   *network.NETWORK
		s   *shell.SHELL
		p   *php.PHP
		bsh *bash.BashInterface
	)

	n, _ = network.Create(cf)
	p = php.Create(cf)
	s = shell.Create(cf)

	bsh = bash.Create(n, s, p)
	modules.Boot(bsh)
	bsh.Start()
}
