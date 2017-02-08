package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		local, remote string
		err           error
		rpc           *razboy.RPCClient
		config        *razboy.Config
		request       *razboy.REQUEST
		response      *razboy.RESPONSE
	)

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Please write the path of the local file to upload")
		return
	}

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	request = razboy.CreateRequest("", config)
	local = os.Args[1]

	if len(os.Args) > 2 {
		remote = os.Args[2]
	} else {
		pathArr := strings.Split(local, "/")
		remote = pathArr[len(pathArr)-1]
	}

	response, err = UploadAction(local, remote, request)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if response.GetResult() != "1" {
		fmt.Fprintln(os.Stderr, "Server havn't upload the file")
		return
	}

	fmt.Fprintln(os.Stdout, response.GetResult())
}

func UploadAction(local, remote string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	request.Action = razboy.CreateUpload(remote)
	request.Upload = true
	request.UploadPath = local

	return razboy.Send(request)
}
