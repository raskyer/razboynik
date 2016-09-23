package common

import (
	"fmt"
	"fuzzer"
	"io"
	"net/http"
	"os"
)

func Download(path, location string) {
	php := fuzzer.PHP.Download(path)
	req, err := fuzzer.NET.Prepare(php)

	if err != nil {
		err.Error()
		return
	}

	resp, err := fuzzer.NET.Send(req)

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
