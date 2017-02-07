package main

import (
	"io"
	"os"

	"fmt"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		local, remote string
		err           error
		rpc           *razboy.RPCClient
		config        *razboy.Config
		request       *razboy.REQUEST
	)

	if len(os.Args) < 1 {
		fmt.Fprintln(os.Stderr, "Please write the path of the file to download")
		return
	}

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	request = razboy.CreateRequest("", config)
	remote = getRemote(os.Args, config.Shellscope)
	local = getLocal(os.Args)

	_, err = DownloadAction(remote, local, request)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: "+err.Error())
		return
	}

	fmt.Fprintln(os.Stdout, "Downloaded successfully to "+local)
}

func getRemote(arr []string, context string) string {
	var path string

	path = arr[1]

	if context != "" {
		path = context + "/" + path
	}

	return path
}

func getLocal(args []string) string {
	if len(args) > 1 {
		return args[1]
	}

	return "output.txt"
}

func DownloadAction(remote, local string, request *razboy.REQUEST) (*razboy.RESPONSE, error) {
	var (
		out *os.File
		res *razboy.RESPONSE
		err error
	)

	request.Action = razboy.CreateDownload(remote)
	res, err = razboy.Send(request)

	if err != nil {
		return res, err
	}

	out, err = os.Create(local)

	defer out.Close()

	if err != nil {
		return nil, err
	}

	_, err = io.Copy(out, res.GetHTTP().Body)

	return res, err
}
