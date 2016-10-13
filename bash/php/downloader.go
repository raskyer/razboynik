package php

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/eatbytes/fuzzcore"
)

func Download(path, location string) {
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
