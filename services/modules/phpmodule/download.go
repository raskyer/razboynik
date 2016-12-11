package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Download(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
	var (
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
		local, remote string
		err           error
	)

	if kc.GetArrLgt() < 2 {
		return kc, errors.New("Please write the path of the file to download")
	}

	request = razboy.CreateRequest("", kc.GetScope(), c)

	remote = _getRemote(kc.GetArr(), kc.GetScope())
	local = kc.GetArrItem(2, "output.txt")

	response, err = DownloadAction(remote, local, request)
	kc.SetResponse(response)

	if err != nil {
		kc.SetResult("Downloaded successfully to " + local)
	}

	return kc, err
}

func DownloadAction(remote, local string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	var (
		out *os.File
		res *razboy.RESPONSE
		err error
	)

	request.Action = phpadapter.CreateDownload(remote)
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
