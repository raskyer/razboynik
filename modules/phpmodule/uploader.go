package phpmodule

import (
	"errors"
	"strings"

	"github.com/eatbytes/fuzz/network"
	"github.com/eatbytes/fuzz/php"
	"github.com/eatbytes/fuzzer/bash"
)

func UploadInit(bc *bash.BashCommand) {
	var path string
	var dir string
	var arr []string
	var err error
	var srv *network.NETWORK
	var ph *php.PHP

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the local file to upload")
		bc.WriteError(err)
		return
	}

	arr = bc.GetArr()
	path = arr[1]
	srv = bc.GetServer()
	ph = bc.GetPHP()

	if bc.GetArrLgt() > 2 {
		dir = arr[2]
	} else {
		pathArr := strings.Split(path, "/")
		lgt := len(pathArr) - 1
		dir = pathArr[lgt]
	}

	bytes, bondary, err := ph.Upload(path, dir)

	if err != nil {
		bc.WriteError(err)
		return
	}

	req, err := srv.PrepareUpload(bytes, bondary)

	if err != nil {
		bc.WriteError(err)
		return
	}

	resp, err := srv.Send(req)

	if err != nil {
		bc.WriteError(err)
		return
	}

	result := srv.GetBodyStr(resp)
	bc.Write(result, nil)
}
