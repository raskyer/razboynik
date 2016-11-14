package shellmodule

import (
	"errors"
	"os"
	"os/exec"

	"github.com/eatbytes/razboynik/bash"
	"github.com/eatbytes/razboynik/modules/phpmodule"
	"github.com/eatbytes/sysgo"
)

func Vim(bc *bash.BashCommand) {
	var (
		remote, local, resp string
		err                 error
		cmd                 *exec.Cmd
	)

	if bc.GetArrLgt() < 2 {
		err = errors.New("Please write the path of the file to edit")
		bc.WriteError(err)
		return
	}

	local = "tmp.txt"
	remote = bc.GetArrItem(1, "")

	_, err = phpmodule.Download(bc.GetServer(), bc.GetPHP(), remote, local)

	if err != nil {
		bc.WriteError(err)
		return
	}

	cmd = exec.Command("vim", local)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err = cmd.Run()

	if err != nil {
		bc.WriteError(err)
		return
	}

	_, err = phpmodule.Upload(bc.GetServer(), bc.GetPHP(), local, remote)

	if err != nil {
		bc.WriteError(err)
		return
	}

	resp, err = sysgo.Call("rm tmp.txt")

	bc.Write(resp, err)
}
