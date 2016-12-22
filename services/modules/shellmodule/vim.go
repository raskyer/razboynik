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

func (vim *Vimcmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		remote, local string
		resp          string
		err           error
		cmd           *exec.Cmd
		request       *razboy.REQUEST
	)

	args := kl.GetArr()

	if len(args) < 1 {
		return vim, errors.New("Please write the path of the file to edit")
	}

	request = razboy.CreateRequest(kl.GetRaw(), config)

	remote = args[0]
	local = "/tmp/tmp-razboynik." + filepath.Ext(remote)

	_, err = phpmodule.DownloadAction(remote, local, request)

	if err != nil {
		return vim, err
	}

	cmd = exec.Command("vim", local)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		return vim, err
	}

	_, err = phpmodule.UploadAction(local, remote, request)

	if err != nil {
		return vim, err
	}

	resp, err = sysgo.Call("rm " + local)
	kl.Write(err, resp)

	return vim, err
}

func (vim *Vimcmd) GetName() string {
	return "vim"
}

func (vim *Vimcmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return nil, true
}

func (vim *Vimcmd) GetResult() []byte {
	return make([]byte, 0)
}

func (vim *Vimcmd) GetResultStr() string {
	return ""
}
