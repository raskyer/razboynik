package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"fmt"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/sysgo"
)

func main() {
	var (
		remote, local string
		resp          string
		err           error
		rpc           *razboy.RPCClient
		config        *razboy.Config
		cmd           *exec.Cmd
		request       *razboy.REQUEST
	)

	rpc = razboy.CreateRPCClient()
	config, err = rpc.GetConfig()

	if len(os.Args) < 1 {
		return
	}

	request = razboy.CreateRequest("vim "+strings.Join(os.Args[1:], " "), config)

	fmt.Println(request)

	remote = os.Args[0]
	local = "/tmp/tmp-razboynik." + filepath.Ext(remote)

	//_, err = phpmodule.DownloadAction(remote, local, request)

	if err != nil {
		return
	}

	cmd = exec.Command("vim", local)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		return
	}

	//_, err = phpmodule.UploadAction(local, remote, request)

	if err != nil {
		return
	}

	resp, err = sysgo.Call("rm " + local)

	fmt.Fprintln(os.Stdout, resp)
}
