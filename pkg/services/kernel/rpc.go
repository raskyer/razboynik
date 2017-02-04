package kernel

import (
	"github.com/eatbytes/razboy"
	"github.com/smallnest/rpcx"
)

type RPCKernel struct {
	Config *razboy.Config
}
type Args interface{}
type Reply interface{}

func (r *RPCKernel) GetConfig(args *Args, reply *Reply) error {
	*reply = r.Config
	return nil
}

func (r *RPCKernel) UpdateConfig(args *razboy.Config, reply *Reply) error {
	*r.Config = *args
	return nil
}

func (r *RPCKernel) UpdatePrompt(args *string, reply *Reply) error {
	k := Boot()
	k.UpdatePrompt(r.Config.Url, *args)

	return nil
}

func CreateRPCServer() *RPCKernel {
	return &RPCKernel{}
}

func StartServer(kernel *RPCKernel) {
	server := rpcx.NewServer()
	server.RegisterName("Kernel", kernel)
	server.Serve("tcp", ":8972")
}
