package bash

import (
	"fmt"
	"fuzzer"
	"fuzzer/src/common"
	"strings"
)

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
	loc := "output.txt"

	if len(arr) > 2 {
		loc = arr[2]
	}

	context := fuzzer.CMD.GetContext()

	if context != "" {
		context = context + "/"
	}

	path = context + path

	common.Download(path, loc)
}
