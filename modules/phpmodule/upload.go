package phpmodule

import (
	"bytes"
	"errors"
	"strings"

	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/php"
	"github.com/eatbytes/razboynik/bash"
)

func UploadInit(bc *bash.BashCommand) {
	var (
		path, dir, resp string
		arr             []string
		err             error
		n               *network.NETWORK
		p               *php.PHP
	)

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the local file to upload")
		bc.WriteError(err)
		return
	}

	arr = bc.GetArr()
	path = arr[1]
	dir = getDir(arr, path)

	n = bc.GetServer()
	p = bc.GetPHP()
	resp, err = Upload(n, p, path, dir)

	bc.Write(resp, err)
}

func getDir(arr []string, path string) string {
	if len(arr) > 2 {
		return arr[2]
	}

	pathArr := strings.Split(path, "/")
	dir := pathArr[len(pathArr)-1]

	return dir
}

func Upload(n *network.NETWORK, p *php.PHP, local, remote string) (string, error) {
	var (
		bondary, body string
		bytes         *bytes.Buffer
		err           error
		req           *network.Request
		resp          *network.Response
	)

	bytes, bondary, err = p.Upload(local, remote)

	if err != nil {
		return "", err
	}

	req, err = n.PrepareUpload(bytes, bondary)

	if err != nil {
		return "", err
	}

	resp, err = n.Send(req)

	if err != nil {
		return "", err
	}

	body = resp.GetResult()

	if body != "1" {
		err = errors.New("Server havn't upload the file")
		return "", err
	}

	return resp.GetBodyStr(), nil
}
