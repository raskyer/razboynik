package bash

import (
	"fmt"
	"strings"

	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/fuzzer/bash/modules"
	"github.com/eatbytes/fuzzer/bash/networking"
	"github.com/eatbytes/fuzzer/bash/parser"
)

func (b *BashInterface) SendUpload(str string) {
	arr := parser.Parse(str)

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

	modules.Upload(path, dir)
}

func (b *BashInterface) SendDownload(str string) {
	arr := parser.Parse(str)

	if arr == nil {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := arr[0]
	loc := "output.txt"

	if len(arr) > 1 {
		loc = arr[1]
	}

	context := fuzzcore.CMD.GetContext()

	if context != "" {
		context = context + "/"
	}

	path = context + path

	modules.Download(path, loc)
}

func (b *BashInterface) SendRawPHP(str string) {
	str, err := parser.ParseStr(str)

	if err != nil {
		return
	}

	raw := fuzzcore.PHP.Raw(str)
	result, err := networking.Process(raw)

	if err != nil {
		err.Error()
		return
	}

	fmt.Println(result)
}
