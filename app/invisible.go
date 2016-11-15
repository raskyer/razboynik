package app

import (
	"errors"
	"fmt"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/normalizer"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

func (app *AppInterface) Invisible(c *cli.Context) {
	var (
		url, referer string
		err          error
		cf           *core.Config
		n            *network.NETWORK
		req          *network.Request
		res          *network.Response
	)

	url = c.Args().First()
	referer = c.Args().Get(1)

	if url == "" || referer == "" {
		err = errors.New("Arguments url and referer are required")
		services.PrintError(err)
		return
	}

	services.PrintSection("Invisible", "Send invisible request")

	cf = &core.Config{
		Url:       url,
		Method:    "GET",
		Parameter: "",
		Shmethod:  0,
		Key:       "",
		Raw:       true,
		Crypt:     false,
	}

	n, err = network.Create(cf)

	if err != nil {
		services.PrintError(err)
		return
	}

	req, err = n.Prepare("")

	if err != nil {
		services.PrintError(err)
		return
	}

	req.Http.Header.Set("Referer", normalizer.Encode(referer))
	res, err = n.Send(req)

	fmt.Println(req.Http)
	fmt.Println(res.Http)

	if err != nil {
		services.PrintError(err)
		return
	}

	services.PrintSection("Result", res.GetBodyStr())
}
