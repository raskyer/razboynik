package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Download(kc *kernel.KernelCmd, request *core.REQUEST) (*kernel.KernelCmd, error) {
	var (
		out           *os.File
		rzRes         *razboy.RazResponse
		local, remote string
		err           error
	)

	if kc.GetArrLgt() < 2 {
		return kc, errors.New("Please write the path of the file to download")
	}

	remote = _getRemote(kc.GetArr(), kc.GetScope())
	local = kc.GetArrItem(2, "output.txt")

	request.Type = "PHP"
	request.Action = phpadapter.CreateDownload(remote, request.PHPc)

	rzRes, err = razboy.Send(request)
	kc.SetResult(rzRes)

	if err != nil {
		return kc, err
	}

	out, err = os.Create(local)

	defer out.Close()

	if err != nil {
		return kc, err
	}

	_, err = io.Copy(out, rzRes.GetHTTP().Body)
	kc.SetBody("Downloaded successfully to " + local)

	return kc, err
}

func _getRemote(arr []string, context string) string {
	var path string

	path = arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}
