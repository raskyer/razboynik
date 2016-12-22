package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Uploadcmd struct{}

func (u *Uploadcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		local, remote string
		err           error
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	args := kl.GetArr()

	if len(args) < 1 {
		return u, errors.New("Please write the path of the local file to upload")
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
		return u, err
	}

	if response.GetResult() != "1" {
		return u, errors.New("Server havn't upload the file")
	}

	return u, nil
}

func (u *Uploadcmd) GetName() string {
	return "-upload"
}

func (u *Uploadcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, false
}

func (u *Uploadcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (u *Uploadcmd) GetResultStr() string {
	return ""
}

func UploadAction(local, remote string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	request.Action = phpadapter.CreateUpload(remote)
	request.Upload = true
	request.UploadPath = local

	return razboy.Send(request)
}
