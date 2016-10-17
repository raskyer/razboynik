package phpmodule

import (
	"fmt"
	"strings"
)

func UploadInit(cmd *BashCommand) {
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

func SendUpload(path, dir string) {
	bytes, bondary, err := fuzzcore.PHP.Upload(path, dir)

	if err != nil {
		err.Error()
		return
	}

	req, err := fuzzcore.NET.PrepareUpload(bytes, bondary)

	if err != nil {
		err.Error()
		return
	}

	resp, err := fuzzcore.NET.Send(req)

	if err != nil {
		err.Error()
		return
	}

	result := fuzzcore.NET.GetBodyStr(resp)
	fmt.Println(result)
}
