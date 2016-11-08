package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/php"
	"github.com/eatbytes/razboynik/bash"
)

func UploadInit(bc *bash.BashCommand) {
	var (
		path, dir string
		arr       []string
		err       error
		n         *network.NETWORK
		p         *php.PHP
	)

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the local file to upload")
		bc.WriteError(err)
		return
	}

	n = bc.GetServer()
	p = bc.GetPHP()

	arr = bc.GetArr()
	path = arr[1]

	if bc.GetArrLgt() > 2 {
		dir = arr[2]
	} else {
		pathArr := strings.Split(path, "/")
		lgt := len(pathArr) - 1
		dir = pathArr[lgt]
	}

	bytes, bondary, err := p.Upload(path, dir)

	if err != nil {
		bc.WriteError(err)
		return
	}

	req, err := n.PrepareUpload(bytes, bondary)

	if err != nil {
		bc.WriteError(err)
		return
	}

	resp, err := n.Send(req)
	body := resp.GetResult()

	if err != nil {
		bc.WriteError(err)
		return
	}

	if body == "1" {
		err = errors.New("Server havn't upload the file")
		bc.WriteError(err)
		return
	}

	bc.Write(resp.GetBodyStr(), nil)
}
