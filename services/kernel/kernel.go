package kernel

import (
	"errors"
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy"
)

type CompleteFunction func(string, *razboy.Config) []string
type KernelCommand interface {
	Init(stdout string, stderr string)
	Exec(*KernelLine, *razboy.Config) (KernelCommand, int, error)
	Write(error, ...interface{}) error
	WriteSuccess(...interface{}) error
	WriteError(error) error
	GetName() string
	GetCompleter() (CompleteFunction, bool)
	GetResultStr() string
	GetResult() []byte
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

func (k *Kernel) Exec(line string, config *razboy.Config) (KernelCommand, int, error) {
	kl := CreateLine(line)

	for _, cmd := range k.commands {
		if cmd.GetName() == kl.name {
			cmd.Init(kl.GetStdout(), kl.GetStderr())

			return cmd.Exec(kl, config)
		}
	}

	if k.def != nil {
		k.def.Init(kl.GetStdout(), kl.GetStderr())
		return k.def.Exec(kl, config)
	}

	return k.Default(kl, config)
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

		_, _, err = k.Exec(line, config)

		if err != nil {
			k.WriteError("&2", err)
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
		err           error
	)

	autocompleter = readline.NewPrefixCompleter()

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

func (k Kernel) Default(kl *KernelLine, config *razboy.Config) (KernelCommand, int, error) {
	return nil, 1, errors.New("No default fonction defined")
}

func dynamicAdapter(completer CompleteFunction, c *razboy.Config) func(string) []string {
	return func(line string) []string {
		return completer(line, c)
	}
}
