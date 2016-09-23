package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/common"
	"strings"
)

func (b *BashInterface) SendUpload(str string) {
	arr := Parse(str)

	if arr == nil {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr[0]
	pathArr := strings.Split(path, "/")
	lgt := len(pathArr) - 1
	dir := pathArr[lgt]

	if len(arr) > 1 {
		dir = arr[1]
	}

	common.Upload(path, dir)
}

func (b *BashInterface) SendDownload(str string) {
	arr := Parse(str)

	if arr == nil {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr[0]
	loc := "output.txt"

	if len(arr) > 1 {
		loc = arr[1]
	}

	context := fuzzer.CMD.GetContext()

	if context != "" {
		context = context + "/"
	}

	path = context + path

	common.Download(path, loc)
}

func (b *BashInterface) SendRawPHP(str string) {
	str, err := ParseStr(str)

	if err != nil {
		return
	}

	raw := fuzzer.PHP.Raw(str)
	result, err := common.Process(raw)

	if err != nil {
		err.Error()
		return
	}

	fmt.Println(result)
}
