package rpc

import (
	"github.com/eatbytes/razboy"
	"github.com/smallnest/rpcx"
)

type RPCKernel struct {
	Config *razboy.Config
}

type Args struct{}
type Reply interface{}

func (r *RPCKernel) GetConfig(args *Args, reply *Reply) error {
	*reply = r.Config
	return nil
}

func CreateRPCKernel(config *razboy.Config) *RPCKernel {
	return &RPCKernel{Config: config}
}

func RPCStart(kernel *RPCKernel) {
	server := rpcx.NewServer()
	server.RegisterName("Kernel", kernel)
	server.Serve("tcp", ":8972")
}
