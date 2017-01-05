package provider

import (
	"net/rpc"
	"os"
	"runtime"

	"net/rpc/jsonrpc"

	"github.com/eatbytes/razboy"
	"github.com/natefinch/pie"
)

const DIR = "providers"

const GOB_CODEC = 0
const JSON_CODEC = 1

const PRINT_CODE = 0
const ASK_CODE = 1

const EXEC_FN = "Plugin.Exec"
const COMPLETER_FN = "Plugin.Completer"

type Args struct {
	Line   string
	Config *razboy.Config
}

type Ask struct {
	RequestName  string
	WaitResponse int
}

type Response struct {
	Code    int
	Content string
	Items   []string
	Ask     Ask
}

type Info struct {
	Path   string
	Name   string
	Method string
}

type Provider interface {
	Exec(*Args, *Response) error
	Completer(*Args, *Response) error
}

func CreateProvider(pi Provider) (*pie.Server, error) {
	var (
		provider pie.Server
		err      error
	)

	provider = pie.NewProvider()
	err = provider.RegisterName("Plugin", pi)

	return &provider, err
}

func GetCodec(name string) int {
	return GOB_CODEC
}

func CallProvider(i *Info, args *Args) (*Response, error) {
	var (
		client *rpc.Client
		resp   *Response
		codec  int
		err    error
	)

	resp = new(Response)

	if runtime.GOOS == "windows" {
		i.Name += ".exe"
	}

	codec = GetCodec(i.Path + i.Name)

	if codec == GOB_CODEC {
		client, err = pie.StartProvider(os.Stderr, i.Path+i.Name)
	} else {
		client, err = pie.StartProviderCodec(jsonrpc.NewClientCodec, os.Stderr, i.Path+i.Name)
	}

	if err != nil {
		return nil, err
	}

	defer client.Close()

	err = client.Call(i.Method, args, resp)

	return resp, err
}
