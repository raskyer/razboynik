package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/reader"
	"io"
	"net/http"
	"os"
	"strings"
)

func (b *BashInterface) Exit(str string) {
	b.Stop()
}

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzer.CMD.Raw(str)
	req, err := fuzzer.NET.Prepare(raw)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	reader.ReadEncode(result)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	req, err := fuzzer.NET.Prepare(cd)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)
	b.ReceiveCd(result)
}

func (b *BashInterface) ReceiveCd(result string) {
	body := fuzzer.Decode(result)
	line := strings.TrimSpace(body)

	if line != "" {
		fuzzer.CMD.SetContext(line)
		b.SetPrompt("\033[31m»\033[0m [Bash]:" + line + "$ ")
		fmt.Println(body)
	}
}

func (b *BashInterface) SendUpload(str string) {
	arr := strings.Fields(str)

	if len(arr) < 2 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr[1]
	pathArr := strings.Split(path, "/")
	lgt := len(pathArr) - 1
	dir := pathArr[lgt]

	if len(arr) > 2 {
		dir = arr[2]
	}

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
	b.ReceiveUpload(result)
}

func (b *BashInterface) ReceiveUpload(result string) {
	if result == "1" {
		fmt.Println("File succeedly upload")
		return
	}

	fmt.Println("An error occured")
}

func (b *BashInterface) SendDownload(str string) {
	php := fuzzer.Download()
	req, err := fuzzer.NET.Prepare(php)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	b.ReceiveDownload(resp)
}

func (b *BashInterface) ReceiveDownload(resp *http.Response) {
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
