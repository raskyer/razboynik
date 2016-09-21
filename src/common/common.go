package common

import (
	"fmt"
	"fuzzer"
	"io"
	"net/http"
	"os"
)

func Upload(path, dir string) {
	bytes, bondary, err := fuzzer.Upload(path, dir)

	if err {
		return
	}

	req, err := fuzzer.NET.PrepareUpload(bytes, bondary)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	ReceiveOne(result, "File upload successfully")
}

func Download(path string) {
	php := fuzzer.Download(path)
	req, err := fuzzer.NET.Prepare(php)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	ReceiveDownload(resp)
}

func ReceiveDownload(resp *http.Response) {
	out, err := os.Create("output.txt")
	defer out.Close()

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded successfully")
}

func ReceiveOne(r, msg string) {
	if r == "1" {
		fmt.Println(msg)
		return
	}

	fmt.Println("An error occured")
}
