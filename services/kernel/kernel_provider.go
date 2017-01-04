package kernel

import (
	"net/rpc"
	"os"
	"runtime"

	"github.com/eatbytes/razboy"
	"github.com/natefinch/pie"
)

const DIR_PROVIDERS = "plugin_provider"

type KernelExternalArgs struct {
	Line   string
	Config *razboy.Config
}

type KernelExternalResponse struct {
	Code     int
	Response string
	Items    []string
}

type KernelExternalCmd interface {
	Exec(*KernelExternalArgs, *KernelExternalResponse) error
	Completer(*KernelExternalArgs, *KernelExternalResponse) error
}

func CreateProvider(kec KernelExternalCmd) (*pie.Server, error) {
	var (
		provider pie.Server
		err      error
	)

	provider = pie.NewProvider()
	err = provider.RegisterName("Plugin", kec)

	return &provider, err
}

func ExecuteProvider(path, method string, args *KernelExternalArgs) (*KernelExternalResponse, error) {
	var (
		client   *rpc.Client
		response *KernelExternalResponse
		err      error
	)

	response = new(KernelExternalResponse)

	if runtime.GOOS == "windows" {
		path = path + ".exe"
	}

	client, err = pie.StartProvider(os.Stderr, path)

	if err != nil {
		return nil, err
	}

	defer client.Close()

	err = client.Call("Plugin."+method, args, response)

	return response, err
}
