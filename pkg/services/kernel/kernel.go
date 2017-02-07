package kernel

import (
	"errors"
	"io"
	"log"
	"os/exec"
	"sort"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/pkg/services/worker/config"
	"github.com/eatbytes/razboynik/pkg/services/worker/printer"
	"github.com/eatbytes/razboynik/pkg/services/worker/usr"
)

type CompleterFunction func(string, *razboy.Config) []string
type GetCompleterFunction func() (CompleterFunction, bool)

type Kernel struct {
	def       string
	commands  map[string]bool
	run       bool
	readline  *readline.Instance
	completer *readline.PrefixCompleter
	rpc       *RPCServer
}

var kInstance *Kernel

func Boot() *Kernel {
	if kInstance == nil {
		kInstance = new(Kernel)
		kInstance.Build()
	}

	return kInstance
}

func (k *Kernel) Build() {
	var (
		cmd  map[string]bool
		dir  []string
		path string
		err  error
	)

	path, err = config.GetPluginPath()

	if err != nil || path == "" {
		path = "../plugin/bin"
		printer.PrintError(errors.New("Can't load configuration. Plugin path will be set : ./plugin/bin"))
	}

	dir = usr.ListDir(path)

	cmd = make(map[string]bool)
	for _, v := range dir {
		cmd[v] = true
	}
	k.SetCommands(cmd)

	k.rpc = CreateRPCServer()
	k.def = "sh"
	go StartServer(k.rpc)
}

func (k *Kernel) Exec(line string, config *razboy.Config) error {
	k.rpc.Config = config
	return k.ExecKernelLine(CreateLine(line), config)
}

func (k *Kernel) ExecKernelLine(l *Line, config *razboy.Config) error {
	if l.GetName() == "exit" {
		k.StopRun()
	}

	if !k.commands[l.GetName()] {
		l = CreateLine("sh " + l.GetName())
	}

	return k.ExecCmd(l)
}

func (k *Kernel) ExecCmd(l *Line) error {
	var cmd *exec.Cmd

	cmd = exec.Command("../plugin/bin/"+l.GetName(), l.GetArg()...)
	cmd.Stdout = l.GetStdout()
	cmd.Stderr = l.GetStderr()
	err := cmd.Run()

	return err
}

func (k *Kernel) Run(config *razboy.Config) error {
	err := k.initReadline(config)

	if err != nil {
		return err
	}

	k.StartRun()

	return k.Loop(config)
}

func (k *Kernel) Loop(config *razboy.Config) error {
	var (
		line string
		err  error
	)

	defer k.readline.Close()
	log.SetOutput(k.readline.Stderr())

	for k.run {
		line, err = k.readline.Readline()

		if err == readline.ErrInterrupt || err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		if len(line) == 0 {
			continue
		}

		err = k.Exec(line, config)

		if err != nil {
			printer.PrintError(err)
		}
	}

	return nil
}

func (k *Kernel) initReadline(c *razboy.Config) error {
	var (
		rl            *readline.Instance
		config        *readline.Config
		autocompleter *readline.PrefixCompleter
		child         *readline.PrefixCompleter
		children      []readline.PrefixCompleterInterface
		cmd           []string
		err           error
	)

	autocompleter = readline.NewPrefixCompleter()

	for key, _ := range k.GetCommands() {
		cmd = append(cmd, key)
	}

	sort.Strings(cmd)

	for _, item := range cmd {
		child = readline.PcItem(item)
		children = append(children, child)
	}

	child = readline.PcItem("exit")
	children = append(children, child)
	// 	if item.Completer == nil {
	// 		child = readline.PcItem(item.Name)
	// 		children = append(children, child)
	// 		continue
	// 	}

	// 	completer, multilevel := item.Completer()

	// 	child = readline.PcItem(
	// 		item.Name,
	// 		readline.PcItemDynamic(dynamicAdapter(completer, c)),
	// 	)

	// 	if multilevel {
	// 		child.MultiLevel = true
	// 	}

	// 	children = append(children, child)
	// }

	autocompleter.SetChildren(children)
	k.completer = autocompleter

	config = &readline.Config{
		Prompt:          "(" + c.Url + ")$ ",
		HistoryFile:     "/tmp/razboynik.tmp",
		AutoComplete:    autocompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	rl, err = readline.NewEx(config)
	k.readline = rl

	return err
}

func (k *Kernel) SetDefault(d string) {
	k.def = d
}

func (k *Kernel) StartRun() {
	k.run = true
}

func (k *Kernel) StopRun() {
	k.run = false
}

func (k *Kernel) GetCommands() map[string]bool {
	return k.commands
}

func (k *Kernel) SetCommands(items map[string]bool) {
	k.commands = items
}

func (k *Kernel) UpdatePrompt(url, scope string) {
	if k.readline == nil {
		return
	}

	k.readline.SetPrompt("(" + url + "):" + scope + "$ ")
}

func dynamicAdapter(completer CompleterFunction, c *razboy.Config) func(string) []string {
	return func(line string) []string {
		return completer(line, c)
	}
}
