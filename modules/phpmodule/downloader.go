package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/razboy/network"
	"github.com/eatbytes/razboy/php"
	"github.com/eatbytes/razboy/shell"
	"github.com/eatbytes/razboynik/bash"
)

func DownloadInit(bc *bash.BashCommand) {
	var (
		path, req string
		err       error
		resp      *network.Response
		n         *network.NETWORK
		s         *shell.SHELL
		p         *php.PHP
	)

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the file to download")
		bc.WriteError(err)
		return
	}

	n, s, p = bc.GetObjects()
	path = getPath(bc.GetArr(), s.GetContext())
	req = p.Download(path)
	resp, err = n.PrepareSend(req)

	if err != nil {
		bc.WriteError(err)
		return
	}

	ReadDownload(resp, bc)
}

func getPath(arr []string, context string) string {
	var path string

	path = arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}

func ReadDownload(resp *network.Response, bc *bash.BashCommand) {
	var (
		location string
		err      error
		out      *os.File
	)

	location = bc.GetArrItem(2, "output.txt")
	out, err = os.Create(location)

	defer out.Close()

	if err != nil {
		bc.WriteError(err)
		return
	}

	_, err = io.Copy(out, resp.Http.Body)

	if err != nil {
		bc.WriteError(err)
		return
	}

	bc.Write("Downloaded successfully to "+location, nil)
}
