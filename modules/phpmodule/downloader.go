package phpmodule

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (b *BashInterface) DownloadInit(bc *BashCommand) {
	if len(bc.arr) < 2 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := bc.arr[1]
	loc := "output.txt"

	if len(bc.arr) > 3 {
		loc = bc.arr[2]
	}

	context := fuzzcore.CMD.GetContext()
	if context != "" {
		context = context + "/"
	}

	path = context + path

	SendDownload(path, loc)
}

func SendDownload(path, location string) {
	php := fuzzcore.PHP.Download(path)
	req, err := fuzzcore.NET.Prepare(php)

	if err != nil {
		err.Error()
		return
	}

	resp, err := fuzzcore.NET.Send(req)

	if err != nil {
		err.Error()
		return
	}

	ReadDownload(resp, location)
}

func ReadDownload(resp *http.Response, location string) {
	out, err := os.Create(location)
	defer out.Close()

	if err != nil {
		err.Error()
		return
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		err.Error()
		return
	}

	fmt.Println("Downloaded successfully those byte: ")
	fmt.Println(resp.Body)
}
