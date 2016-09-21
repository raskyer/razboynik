package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/commander"
	"fuzzer/src/common"
	"fuzzer/src/reader"
	"strings"
)

func (b *BashInterface) SendRaw(str string) {
	raw := fuzzer.CMD.Raw(str)
	result, err := commander.Process(raw)

	if err {
		return
	}

	reader.ReadEncode(result)
}

func (b *BashInterface) SendCd(str string) {
	cd := fuzzer.CMD.Cd(str)
	result, err := commander.Process(cd)

	if err {
		return
	}

	b.ReceiveCd(result)
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

	common.Upload(path, dir)
}

func (b *BashInterface) SendDownload(str string) {
	arr := strings.Fields(str)

	if len(arr) < 2 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr[1]

	common.Download(path)
}
