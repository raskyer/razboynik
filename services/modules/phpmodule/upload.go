package phpmodule

import (
	"bytes"
	"errors"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Upload(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request       *razboy.REQUEST
		rzRes         *razboy.RazResponse
		local, remote string
		arr           []string
		err           error
	)

	if kc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the local file to upload")
		return kc, err
	}

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

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

func UploadAction(local, remote string, request *razboy.REQUEST) (*razboy.RazResponse, error) {
	var (
		upload string
		data   *bytes.Buffer
		err    error
	)

	upload, data, err = phpadapter.CreateUpload(local, remote, request.Raw)

	if err != nil {
		return nil, err
	}

	request.Action = upload
	request.Upload = true
	request.Buffer = data

	return razboy.Send(request)
}
