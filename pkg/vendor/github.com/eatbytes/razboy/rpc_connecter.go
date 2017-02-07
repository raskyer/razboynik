package razboy

import (
	"github.com/smallnest/rpcx"
)

//Const
const (
	OBJECT    = "Kernel"
	PROTOCOL  = "tcp"
	ADDR      = ":8972"
	GETCONFIG = "Kernel.GetConfig"
	SETCONFIG = "Kernel.SetConfig"
	SETPROMPT = "Kernel.SetPrompt"

	RPCERROR     int = 1
	NETWORKERROR int = 2
)

type RPCClient struct {
	client *rpcx.Client
}

type Args interface{}
type Reply interface{}

//CreateRPCClient Create a new client for RPC
func CreateRPCClient() *RPCClient {
	selector := &rpcx.DirectClientSelector{Network: PROTOCOL, Address: ADDR}
	client := rpcx.NewClient(selector)

	return &RPCClient{client}
}

//GetConfig Get the config object from rpc
func (r *RPCClient) GetConfig() (*Config, error) {
	config := new(Config)
	err := r.client.Call(GETCONFIG, nil, config)

	return config, err
}

//SetConfig Set the config object to rpc
func (r *RPCClient) SetConfig(config *Config) error {
	return r.client.Call(SETCONFIG, config, nil)
}

//SetPrompt Set the prompt to rpc
func (r *RPCClient) SetPrompt(url, path string) error {
	return r.client.Call(SETPROMPT, &[]string{url, path}, nil)
}
