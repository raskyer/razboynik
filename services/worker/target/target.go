package target

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/fatih/color"
)

type Target struct {
	Name   string         `json:"name"`
	Config *razboy.Config `json:"config"`
}

func _getInput(txt, def string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target's " + color.YellowString(txt) + " (\"" + color.MagentaString(def) + "\"): ")

	tmp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)

		return _getInput(txt, def)
	}

	tmp = strings.TrimSpace(tmp)

	if tmp == "" {
		tmp = def
	} else if tmp == "x" {
		tmp = ""
	}

	color.Green(tmp)

	return tmp
}

func CreateTarget() *Target {
	var (
		target *Target
	)

	target = new(Target)
	target.Config = razboy.NewConfig()
	EditTarget(target)

	return target
}

func EditTarget(target *Target) {
	var tmp string

	target.Name = _getInput("name", target.Name)
	target.Config.Url = _getInput("URL", target.Config.Url)

	tmp = _getInput("method ['GET', 'POST', 'HEADER', 'COOKIE']", razboy.MethodToStr(target.Config.Method))
	target.Config.Method = razboy.MethodToInt(tmp)

	target.Config.Parameter = _getInput("parameter", target.Config.Parameter)

	tmp = _getInput("shell method ['system', 'shell_exec', 'proc_open', 'passthru']", razboy.ShellmethodToStr(target.Config.Shellmethod))
	target.Config.Shellmethod = razboy.ShellmethodToInt(tmp)

	target.Config.Shellscope = _getInput("shell scope ['./', '/']", target.Config.Shellscope)
	target.Config.Key = _getInput("key", target.Config.Key)
}
