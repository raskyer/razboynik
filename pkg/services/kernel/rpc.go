package kernel

import (
	"github.com/eatbytes/razboy"
	"github.com/smallnest/rpcx"
)

type RPCServer struct {
	Config *razboy.Config
}

type Args interface{}
type Reply interface{}

func CreateRPCServer() *RPCServer {
	return &RPCServer{}
}

func StartServer(r *RPCServer) {
	server := rpcx.NewServer()
	server.RegisterName(razboy.OBJECT, r)
	server.Serve(razboy.PROTOCOL, razboy.ADDR)
}

func (r *RPCServer) GetConfig(args *Args, reply *Reply) error {
	*reply = r.Config

	return nil
}

func (r *RPCServer) SetConfig(args *razboy.Config, reply *Reply) error {
	*r.Config = *args

	return nil
}

func (r *RPCServer) SetPrompt(args *[]string, reply *Reply) error {
	k := Boot()
	k.UpdatePrompt((*args)[0], (*args)[1])

	return nil
}
