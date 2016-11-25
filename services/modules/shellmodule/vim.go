package shellmodule

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/config"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/modules/phpmodule"
	"github.com/eatbytes/sysgo"
)

func Vim(kc *kernel.KernelCmd, c *config.Config) (*kernel.KernelCmd, error) {
	var (
		request             *razboy.REQUEST
		remote, local, resp string
		err                 error
		cmd                 *exec.Cmd
	)

	if kc.GetArrLgt() < 2 {
		return kc, errors.New("Please write the path of the file to edit")
	}

	request = razboy.CreateRequest(
		[4]string{c.Url, c.Method, c.Parameter, c.Key},
		[2]string{c.Shellmethod, kc.GetScope()},
		[2]bool{c.Raw, false},
	)

	remote = kc.GetArrItem(1)
	local = "/tmp/tmp-razboynik." + filepath.Ext(remote)

	_, err = phpmodule.DownloadAction(remote, local, request)

	if err != nil {
		return kc, err
	}

	cmd = exec.Command("vim", local)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		return kc, err
	}

	_, err = phpmodule.UploadAction(local, remote, request)

	if err != nil {
		return kc, err
	}

	resp, err = sysgo.Call("rm " + local)
	kc.SetBody(resp)

	return kc, err
}
