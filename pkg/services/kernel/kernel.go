package kernel

import (
	"errors"
	"io"
	"log"
	"os/exec"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/pkg/services/rpc"
	"github.com/eatbytes/razboynik/pkg/services/worker/configuration"
	"github.com/eatbytes/razboynik/pkg/services/worker/printer"
	"github.com/eatbytes/razboynik/pkg/services/worker/usr"
)

type CompleterFunction func(string, *razboy.Config) []string
type GetCompleterFunction func() (CompleterFunction, bool)

type Kernel struct {
	def           string
	commands      []string
	path          string
	readline      *readline.Instance
	autocompleter *readline.PrefixCompleter
	run           bool
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
		config *configuration.Configuration
		dir    []string
		path   string
		err    error
	)

	path = "./plugin/bin"
	config, err = configuration.GetConfiguration()

	if err != nil {
		printer.PrintError(errors.New("Can't load configuration. Plugin path will be set : ./plugin/bin"))
	} else {
		path = config.PluginDir
	}

	dir = usr.ListDir(path)
	k.SetCommands(dir)
}

func (k *Kernel) Exec(line string, config *razboy.Config) error {
	return k.ExecKernelLine(CreateLine(line), config)
}

func (k *Kernel) ExecKernelLine(l *Line, config *razboy.Config) error {
	if l.GetName() == "exit" {
		k.StopRun()
	}

	for _, cmd := range k.commands {
		if cmd == l.GetName() {
			return k.ExecCmd(l)
		}
	}

	if k.def != "" {
		return nil
	}

	return k.Default(l, config)
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

	krpc := rpc.CreateRPCKernel(config)
	go rpc.RPCStart(krpc)

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
		err           error
	)

	autocompleter = readline.NewPrefixCompleter()

	for _, item := range k.GetCommands() {
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

	k.autocompleter = autocompleter

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

func (k *Kernel) Default(l *Line, config *razboy.Config) error {
	return errors.New("No default fonction defined")
}

func (k *Kernel) StartRun() {
	k.run = true
}

func (k *Kernel) StopRun() {
	k.run = false
}

func (k *Kernel) GetCommands() []string {
	return k.commands
}

func (k *Kernel) SetCommands(items []string) {
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
