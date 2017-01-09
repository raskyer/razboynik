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

var Vimitem = kernel.Item{
	Name: "vim",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			remote, local string
			resp          string
			err           error
			cmd           *exec.Cmd
			request       *razboy.REQUEST
		)

		args := l.GetArg()

		if len(args) < 1 {
			return kernel.Response{Err: errors.New("Please write the path of the file to edit")}
		}

		request = razboy.CreateRequest(l.GetRaw(), config)

		remote = args[0]
		local = "/tmp/tmp-razboynik." + filepath.Ext(remote)

		_, err = phpmodule.DownloadAction(remote, local, request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		cmd = exec.Command("vim", local)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		err = cmd.Run()

		if err != nil {
			return kernel.Response{Err: err}
		}

		_, err = phpmodule.UploadAction(local, remote, request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		resp, err = sysgo.Call("rm " + local)
		kernel.WriteSuccess(l.GetStdout(), resp)

		return kernel.Response{Body: resp, Err: err}
	},
}
