package razboy

import (
	"github.com/smallnest/rpcx"
)

//Const
const (
	OBJECT         = "Kernel"
	PROTOCOL       = "tcp"
	ADDR           = ":8972"
	GET_CONFIG     = "Kernel.GetConfig"
	SET_CONFIG     = "Kernel.SetConfig"
	SET_PROMPT     = "Kernel.SetPrompt"
	REQUEST_PLUGIN = "Kernel.RequestPlugin"
	REQUEST_SYSTEM = "Kernel.RequestSystem"

	RPC_ERROR     int = 1
	NETWORK_ERROR int = 2
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
	err := r.client.Call(GET_CONFIG, nil, config)

	return config, err
}

//SetConfig Set the config object to rpc
func (r *RPCClient) SetConfig(config *Config) error {
	return r.client.Call(SET_CONFIG, config, nil)
}

//SetPrompt Set the prompt to rpc
func (r *RPCClient) SetPrompt(url, path string) error {
	return r.client.Call(SET_PROMPT, &[]string{url, path}, nil)
}

//RequestPlugin Request an other plugin
func (r *RPCClient) RequestPlugin(cmd string, reply [][]byte) error {
	return r.client.Call(REQUEST_PLUGIN, &cmd, &reply)
}

//RequestSystem Request an action on host system
func (r *RPCClient) RequestSystem(cmd []string) error {
	return r.client.Call(REQUEST_SYSTEM, &cmd, nil)
}
