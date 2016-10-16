package bash

import (
	"fmt"
	"strings"

	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/fuzzer/bash/modules"
	"github.com/eatbytes/fuzzer/bash/networking"
)

func (b *BashInterface) SendUpload(cmd *BashCommand) {
	if len(cmd.arr) < 2 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := cmd.arr[1]
	var dir string

	if len(cmd.arr) > 2 {
		dir = cmd.arr[3]
	} else {
		pathArr := strings.Split(path, "/")
		lgt := len(pathArr) - 1
		dir = pathArr[lgt]
	}

	modules.Upload(path, dir)
}

func (b *BashInterface) SendDownload(cmd *BashCommand) {
	if len(cmd.arr) < 2 {
		fmt.Println("Please write the path of the local file to upload")
		return
	}

	path := cmd.arr[1]
	loc := "output.txt"

	if len(cmd.arr) > 3 {
		loc = cmd.arr[2]
	}

	context := fuzzcore.CMD.GetContext()
	if context != "" {
		context = context + "/"
	}

	path = context + path

	modules.Download(path, loc)
}

func (b *BashInterface) SendRawPHP(cmd *BashCommand) {
	raw := fuzzcore.PHP.Raw(cmd.str)
	result, err := networking.Process(raw)
	cmd.Write(result, err)
}
