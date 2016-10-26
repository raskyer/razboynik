package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/php"
	"github.com/eatbytes/fuzz/shell"
	"github.com/eatbytes/fuzzer/bash"
)

func DownloadInit(bc *bash.BashCommand) {
	var path string
	var req string
	var err error
	var resp *network.Response
	var srv *network.NETWORK
	var shl *shell.SHELL
	var ph *php.PHP

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the file to download")
		bc.WriteError(err)
		return
	}

	srv, shl, ph = bc.GetObjects()

	path = getPath(bc.GetArr(), shl.GetContext())
	req = ph.Download(path)

	resp, err = srv.PrepareSend(req)

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
	var location string
	location = bc.GetArrItem(2, "output.txt")

	out, err := os.Create(location)
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
