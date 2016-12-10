package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Upload(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request       *razboy.REQUEST
		rzRes         *razboy.RESPONSE
		local, remote string
		arr           []string
		err           error
	)

	if kc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the local file to upload")
		return kc, err
	}

	request = razboy.CreateRequest("", kc.GetScope(), c)

	arr = kc.GetArr()
	local = arr[1]

	if kc.GetArrLgt() > 2 {
		remote = arr[2]
	} else {
		pathArr := strings.Split(local, "/")
		lgt := len(pathArr) - 1
		remote = pathArr[lgt]
	}

	rzRes, err = UploadAction(local, remote, request)
	kc.SetResult(rzRes)

	if err != nil {
		return kc, err
	}

	if kc.GetResult() == "1" {
		return kc, errors.New("Server havn't upload the file")
	}

	return kc, nil
}

func UploadAction(local, remote string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	request.Action = phpadapter.CreateUpload(remote)
	request.Upload = true
	request.UploadPath = local

	return razboy.Send(request)
}
