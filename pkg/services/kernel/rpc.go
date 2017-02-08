package kernel

import (
	"bytes"

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

func (r *RPCServer) RequestOther(args *string, reply *[][]byte) error {
	var (
		k      *Kernel
		l      *Line
		stdout *bytes.Buffer
		stderr *bytes.Buffer
		err    error
	)

	stdout = new(bytes.Buffer)
	stderr = new(bytes.Buffer)

	k = Boot()
	l = CreateLine(*args)

	err = k.ExecCmd(l, stdout, stderr)

	if err != nil {
		return err
	}

	*reply = [][]byte{stdout.Bytes(), stderr.Bytes()}

	return nil
}
