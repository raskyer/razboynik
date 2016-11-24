package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Upload(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		rzRes                 *razboy.RazResponse
		local, remote, upload string
		arr                   []string
		err                   error
	)

	if kc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the local file to upload")
		return kc, err
	}

	arr = kc.GetArr()
	local = arr[1]

	if kc.GetArrLgt() > 2 {
		remote = arr[2]
	} else {
		pathArr := strings.Split(local, "/")
		lgt := len(pathArr) - 1
		remote = pathArr[lgt]
	}

	upload, err = phpadapter.CreateUpload(local, remote, request.PHPc)

	if err != nil {
		return kc, err
	}

	request.Type = "PHP"
	request.Action = upload
	request.PHPc.Upload = true

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	if err != nil {
		return kc, err
	}

	if kc.GetResult() == "1" {
		return kc, errors.New("Server havn't upload the file")
	}

	return kc, nil
}
