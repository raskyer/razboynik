package phpmodule

import (
	"errors"
	"io"
	"os"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/adapter/phpadapter"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Download(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request       *razboy.REQUEST
		rzRes         *razboy.RazResponse
		local, remote string
		err           error
	)

	if kc.GetArrLgt() < 2 {
		return kc, errors.New("Please write the path of the file to download")
	}

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

	remote = _getRemote(kc.GetArr(), kc.GetScope())
	local = kc.GetArrItem(2, "output.txt")

	rzRes, err = DownloadAction(remote, local, request)
	kc.SetResult(rzRes)
	kc.SetBody("Downloaded successfully to " + local)

	return kc, err
}

func DownloadAction(remote, local string, request *razboy.REQUEST) (*razboy.RazResponse, error) {
	var (
		out   *os.File
		rzRes *razboy.RazResponse
		err   error
	)

	request.Action = phpadapter.CreateDownload(remote, false)
	rzRes, err = razboy.Send(request)

	if err != nil {
		return rzRes, err
	}

	out, err = os.Create(local)

	defer out.Close()

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(out, rzRes.GetHTTP().Body)

	return rzRes, err
}

func _getRemote(arr []string, context string) string {
	var path string

	path = arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}
