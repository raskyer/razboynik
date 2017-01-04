package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Downloadcmd struct{}

func (d *Downloadcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		local, remote string
		err           error
		request       *razboy.REQUEST
	)

	args := kl.GetArr()

	if len(args) < 1 {
		return errors.New("Please write the path of the file to download")
	}

	request = razboy.CreateRequest(kl.GetRaw(), config)
	remote = _getRemote(args, config.Shellscope)

	if len(args) > 1 {
		local = args[1]
	} else {
		local = "output.txt"
	}

	_, err = DownloadAction(remote, local, request)
	kernel.Write(kl.GetStdout(), kl.GetStderr(), err, "Downloaded successfully to "+local)

	return err
}

func (d *Downloadcmd) GetName() string {
	return "-download"
}

func (d *Downloadcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}

func DownloadAction(remote, local string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	var (
		out *os.File
		res *razboy.RESPONSE
		err error
	)

	request.Action = razboy.CreateDownload(remote)
	res, err = razboy.Send(request)

	if err != nil {
		return res, err
	}

	out, err = os.Create(local)

	defer out.Close()

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(out, res.GetHTTP().Body)

	return res, err
}

func _getRemote(arr []string, context string) string {
	var path string

	path = arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}
