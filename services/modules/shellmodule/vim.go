package shellmodule

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/phpmodule"
	"github.com/eatbytes/sysgo"
)

type Vimcmd struct{}

func (vim *Vimcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
	var (
		remote, local string
		resp          string
		err           error
		cmd           *exec.Cmd
		request       *razboy.REQUEST
	)

	args := kl.GetArr()

	if len(args) < 1 {
		return kernel.KernelResponse{Err: errors.New("Please write the path of the file to edit")}
	}

	request = razboy.CreateRequest(kl.GetRaw(), config)

	remote = args[0]
	local = "/tmp/tmp-razboynik." + filepath.Ext(remote)

	_, err = phpmodule.DownloadAction(remote, local, request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	cmd = exec.Command("vim", local)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	_, err = phpmodule.UploadAction(local, remote, request)

	if err != nil {
		return kernel.KernelResponse{Err: err}
	}

	resp, err = sysgo.Call("rm " + local)
	kernel.WriteSuccess(kl.GetStdout(), resp)

	return kernel.KernelResponse{Body: resp}
}

func (vim *Vimcmd) GetName() string {
	return "vim"
}

func (vim *Vimcmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, true
}
