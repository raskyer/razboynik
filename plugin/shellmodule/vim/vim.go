package main

import (
	"os"
	"os/exec"
	"path/filepath"

	"fmt"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/sysgo"
)

func main() {
	var (
		remote, local string
		resp          string
		err           error
	)

	if len(os.Args) < 2 {
		return
	}

	remote = os.Args[1]
	local = getLocal(remote)

	err = download(remote, local)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	err = vim(local)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	err = upload(local, remote)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	resp, err = sysgo.Call("rm " + local)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Println(resp)
}

func getLocal(remote string) string {
	ext := filepath.Ext(remote)

	if ext == "" {
		ext = "txt"
	}

	return "/tmp/tmp-razboynik." + ext
}

func download(remote, local string) error {
	rpc := razboy.CreateRPCClient()
	reply := make([][]byte, 2)

	err := rpc.RequestOther("download "+remote+" "+local, reply)

	if err != nil {
		return err
	}

	fmt.Println(string(reply[0]))

	return nil
}

func upload(local, remote string) error {
	rpc := razboy.CreateRPCClient()
	reply := make([][]byte, 2)

	err := rpc.RequestOther("upload "+local+" "+remote, reply)

	if err != nil {
		return err
	}

	fmt.Println("Stdout: ", string(reply[0]))
	fmt.Println("Stderr: ", string(reply[1]))

	return nil
}

func vim(local string) error {
	cmd := exec.Command("vim", local)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
