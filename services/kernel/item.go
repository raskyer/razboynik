package kernel

import "github.com/eatbytes/razboy"

type ItemExecFunction func(*Line, *razboy.Config) Response

type RPCInfo struct {
	Addr string
	Port int
}

type Item struct {
	Name      string
	Exec      ItemExecFunction
	Completer GetCompleterFunction
	RPC       *RPCInfo
}
