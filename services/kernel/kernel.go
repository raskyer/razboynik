package kernel

import (
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy"
)

type KernelFunction func(*KernelCmd, *razboy.Config) (*KernelCmd, error)
type CompleteFunction func(string, *razboy.Config) []string

type KernelItem struct {
	Name       string
	Fn         KernelFunction
	Callback   CompleteFunction
	MultiLevel bool
}

type Kernel struct {
	def      *KernelItem
	items    []*KernelItem
	readline *readline.Instance
	former   *KernelCmd
	run      bool
}

var kInstance *Kernel

func Boot(def ...*KernelItem) *Kernel {
	var defaultFn *KernelItem

	if kInstance == nil {
		defaultFn = &KernelItem{
			Name: "kernel.default",
			Fn:   KernelDefault,
		}

		if len(def) > 0 {
			defaultFn = def[0]
		}

		kInstance = &Kernel{
			def: defaultFn,
		}
	}

	return kInstance
}

func (k *Kernel) Exec(kc *KernelCmd, config *razboy.Config) (*KernelCmd, error) {
	for _, item := range k.items {
		if item.Name == kc.name {
			return item.Fn(kc, config)
		}
	}

	return k.def.Fn(kc, config)
}

func (k *Kernel) Run(config *razboy.Config) error {
	var err error

	err = k.initReadline(config)

	if err != nil {
		return err
	}

	k.run = true
	k._loop(config)

	return nil
}

func (k *Kernel) _loop(config *razboy.Config) {
	var (
		kc, fkc *KernelCmd
		line    string
		err     error
	)

	defer k.readline.Close()
	log.SetOutput(k.readline.Stderr())

	for k.run {
		line, err = k.readline.Readline()

		if err == readline.ErrInterrupt || err == io.EOF {
			return
		}

		if len(line) == 0 {
			continue
		}

		if fkc != nil {
			k.SetFormerCmd(fkc)
		}

		kc = CreateCmd(line)
		fkc, err = k.Exec(kc, config)

		if err != nil {
			fkc.WriteError(err)
			continue
		}

		fkc.WriteSuccess(fkc.GetResult())
	}
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

	for _, item := range k.GetItems() {
		if item.Callback != nil {
			child = readline.PcItem(
				item.Name,
				readline.PcItemDynamic(dynamicAdapter(c, item)),
			)

			if item.MultiLevel {
				child.MultiLevel = true
			}
		} else {
			child = readline.PcItem(item.Name)
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

func dynamicAdapter(c *razboy.Config, item *KernelItem) func(string) []string {
	return func(line string) []string {
		return item.Callback(line, c)
	}
}
