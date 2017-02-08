package main

import (
	"os"

	"fmt"

	"github.com/eatbytes/razboy"
)

func main() {
	var (
		rpc *razboy.RPCClient
		err error
	)

	rpc = razboy.CreateRPCClient()
	err = rpc.RequestSystem(os.Args[1:len(os.Args)])

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
