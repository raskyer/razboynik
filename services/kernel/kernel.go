package kernel

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"reflect"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy"
)

type CompleterFunction func(string, *razboy.Config) []string
type KernelCommand interface {
	Exec(*KernelLine, *razboy.Config) error
	GetName() string
	GetCompleter() (CompleterFunction, bool)
}

type Kernel struct {
	def      KernelCommand
	commands []KernelCommand
	readline *readline.Instance
	run      bool
}

var kInstance *Kernel

func Boot() *Kernel {
	if kInstance == nil {
		kInstance = new(Kernel)
	}

	return kInstance
}

func (k *Kernel) Exec(line string, config *razboy.Config) error {
	kl := CreateLine(line)

	if strings.HasPrefix(kl.GetName(), "#") {
		return k.ExecByProvider(kl)
	}

	for _, cmd := range k.commands {
		if cmd.GetName() == kl.GetName() {
			return cmd.Exec(kl, config)
		}
	}

	if k.def != nil {
		return k.def.Exec(kl, config)
	}

	return k.Default(kl, config)
}

func (k *Kernel) ExecByProvider(kl *KernelLine) error {
	args := new(KernelExternalArgs)
	args.Line = kl.GetRaw()

	name := DIR_PROVIDERS + "/" + strings.TrimPrefix(kl.GetName(), "#")
	resp, err := ExecuteProvider(name, "Exec", args)

	if err != nil {
		return WriteError(kl.GetStderr(), err)
	}

	return WriteSuccess(kl.GetStdout(), resp.Response)
}

func (k *Kernel) Run(config *razboy.Config) error {
	var err error

	err = k.initReadline(config)

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
			fmt.Println(err)
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
		filesInfo     []os.FileInfo
		dir           string
		err           error
	)

	autocompleter = readline.NewPrefixCompleter()

	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	dir = dir + "/" + DIR_PROVIDERS + "/"
	filesInfo, err = ioutil.ReadDir(dir)
	if err == nil {
		for _, item := range filesInfo {
			if strings.Contains(item.Name(), ".") {
				continue
			}

			child = readline.PcItem(
				"#"+item.Name(),
				readline.PcItemDynamic(dynamicExternalAdapter(dir+item.Name())),
			)

			autocompleter.SetChildren(append(autocompleter.GetChildren(), child))
		}
	}

	for _, item := range k.GetCommands() {
		completer, multilevel := item.GetCompleter()

		if completer != nil {
			child = readline.PcItem(
				item.GetName(),
				readline.PcItemDynamic(dynamicAdapter(completer, c)),
			)

			if multilevel {
				child.MultiLevel = true
			}
		} else {
			child = readline.PcItem(item.GetName())
		}

		autocompleter.SetChildren(append(autocompleter.GetChildren(), child))
	}

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

func (k Kernel) Default(kl *KernelLine, config *razboy.Config) error {
	return errors.New("No default fonction defined")
}

func (k Kernel) GetCommands() []KernelCommand {
	return k.commands
}

func (k *Kernel) StartRun() {
	k.run = true
}

func (k *Kernel) StopRun() {
	k.run = false
}

func (k *Kernel) UpdatePrompt(url, scope string) {
	if k.readline == nil {
		return
	}

	k.readline.SetPrompt("(" + url + "):" + scope + "$ ")
}

func (k *Kernel) SetDefault(d KernelCommand) {
	k.def = d
}

func (k *Kernel) SetCommands(cmd []KernelCommand) {
	k.commands = cmd
}

func Write(stdout, stderr *os.File, e error, i ...interface{}) error {
	if e != nil {
		return WriteError(stderr, e)
	}

	return WriteSuccess(stdout, i...)
}

func WriteSuccess(stdout *os.File, i ...interface{}) error {
	var e error

	for _, v := range i {
		if reflect.TypeOf(v).Kind() == reflect.String {
			_, e = fmt.Fprint(stdout, strings.TrimSpace(v.(string)))
		} else {
			_, e = fmt.Fprint(stdout, v)
		}

		if e != nil {
			fmt.Println(e)
		}
	}

	fmt.Fprint(stdout, "\n")

	return e
}

func WriteError(stderr *os.File, err error) error {
	_, e := fmt.Fprintln(stderr, err.Error())

	return e
}

func dynamicAdapter(completer CompleterFunction, c *razboy.Config) func(string) []string {
	return func(line string) []string {
		return completer(line, c)
	}
}

func dynamicExternalAdapter(path string) func(string) []string {
	return func(line string) []string {
		args := new(KernelExternalArgs)
		args.Line = line

		ker, err := ExecuteProvider(path, "Completer", args)

		if err != nil {
			fmt.Println(err)
			return []string{}
		}

		return ker.Items
	}
}
