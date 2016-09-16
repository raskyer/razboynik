package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/reader"
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

	fuzzer.NET.Send(req, reader.ReadEncode)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	req, err := fuzzer.NET.Prepare(cd)

	if err {
		return
	}

	fuzzer.NET.Send(req, b.ReceiveCd)
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
	dir := "./titi.txt"

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

	fuzzer.NET.Send(req, b.ReceiveUpload)
}

func (b *BashInterface) ReceiveUpload(result string) {
	if result == "1" {
		fmt.Println("File succeedly upload")
		return
	}

	fmt.Println("An error occured")
}

func (b *BashInterface) SendDownload(str string) {

}
