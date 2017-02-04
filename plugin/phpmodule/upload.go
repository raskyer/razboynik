package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Uploaditem = kernel.Item{
	Name: "-upload",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			local, remote string
			err           error
			request       *razboy.REQUEST
			response      *razboy.RESPONSE
		)

		args := l.GetArg()

		if len(args) < 1 {
			return kernel.Response{Err: errors.New("Please write the path of the local file to upload")}
		}

		request = razboy.CreateRequest(l.GetRaw(), config)
		local = args[0]

		if len(args) > 1 {
			remote = args[1]
		} else {
			pathArr := strings.Split(local, "/")
			lgt := len(pathArr) - 1
			remote = pathArr[lgt]
		}

		response, err = UploadAction(local, remote, request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		if response.GetResult() != "1" {
			return kernel.Response{Err: errors.New("Server havn't upload the file")}
		}

		kernel.WriteSuccess(l.GetStdout(), response.GetResult())

		return kernel.Response{Body: response.GetResult()}
	},
}

func UploadAction(local, remote string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	request.Action = razboy.CreateUpload(remote)
	request.Upload = true
	request.UploadPath = local

	return razboy.Send(request)
}
