package main

import (
	"os"
	"path/filepath"

	"fmt"

	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/sysgo"
)

func main() {
	var (
		remote, local string
		err           error
		rpc           *razboy.RPCClient
	)

	if len(os.Args) < 2 {
		return
	}

	rpc = razboy.CreateRPCClient()
	remote = os.Args[1]
	local = getLocal(remote)

	err = download(remote, local, rpc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	err = vim(local, rpc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	err = upload(local, remote, rpc)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	_, err = sysgo.Call("rm " + local)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Println("Edited successfully")
}

func getLocal(remote string) string {
	ext := filepath.Ext(remote)

	if ext == "" {
		ext = "txt"
	}

	return "/tmp/tmp-razboynik." + ext
}

func download(remote, local string, rpc *razboy.RPCClient) error {
	plugin := "download " + remote + " " + local
	reply := make([][]byte, 2)

	err := rpc.RequestPlugin(plugin, reply)

	if err != nil {
		return err
	}

	fmt.Println("Stdout: ", strings.TrimSpace(string(reply[0])))
	fmt.Println("Stderr: ", strings.TrimSpace(string(reply[1])))

	return nil
}

func upload(local, remote string, rpc *razboy.RPCClient) error {
	plugin := "upload " + local + " " + remote
	reply := make([][]byte, 2)

	err := rpc.RequestPlugin(plugin, reply)

	if err != nil {
		return err
	}

	fmt.Println("Stdout: ", strings.TrimSpace(string(reply[0])))
	fmt.Println("Stderr: ", strings.TrimSpace(string(reply[1])))

	return nil
}

func vim(local string, rpc *razboy.RPCClient) error {
	system := []string{"vim", local}

	return rpc.RequestSystem(system)
}
