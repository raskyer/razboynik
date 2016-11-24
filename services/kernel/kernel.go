package kernel

import (
	"errors"
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy/core"
)

type KernelFunction func(*KernelCmd, *core.REQUEST) (*KernelCmd, error)

type KernelItem struct {
	Name string
	Fn   KernelFunction
}

type Kernel struct {
	def      *KernelItem
	items    []*KernelItem
	readline *readline.Instance
	commons  []string
	run      bool
}

var kInstance *Kernel

func Boot(def ...*KernelItem) *Kernel {
	var defaultFn *KernelItem

	if kInstance == nil {
		defaultFn = &KernelItem{
			Name: "kernel.default",
			Fn:   _kernelDefault,
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

func (k *Kernel) Exec(kc *KernelCmd, request *core.REQUEST) (*KernelCmd, error) {
	for _, item := range k.items {
		if item.Name == kc.name {
			return item.Fn(kc, request)
		}
	}

	return k.def.Fn(kc, request)
}

func (k *Kernel) Run(request *core.REQUEST) error {
	var err error

	err = k._initReadline(request.SRVc.Url)

	if err != nil {
		return err
	}

	k.run = true
	k._loop(request)

	return nil
}

func (k Kernel) GetDefaultItem() *KernelItem {
	return k.def
}

func (k Kernel) GetItems() []*KernelItem {
	return k.items
}

func (k Kernel) GetItemsName() []string {
	var names []string

	for _, item := range k.items {
		names = append(names, item.Name)
	}

	return names
}

func (k Kernel) GetCommons() []string {
	return k.commons
}

func (k *Kernel) SetDefault(item *KernelItem) {
	k.def = item
}

func (k *Kernel) SetItems(items []*KernelItem) {
	k.items = items
}

func (k *Kernel) AddItem(item *KernelItem) {
	k.items = append(k.items, item)
}

func (k *Kernel) Stop() {
	k.run = false
}

func (k *Kernel) UpdatePrompt(url, scope string) {
	if k.readline == nil {
		return
	}

	k.readline.SetPrompt("(" + url + "):" + scope + "$ ")
}

func (k *Kernel) _cleanRequest(request *core.REQUEST) {
	request.Type = ""
	request.Action = ""
}

func (k *Kernel) _loop(request *core.REQUEST) {
	var (
		kc          *KernelCmd
		fkc         *KernelCmd
		line, scope string
		err         error
	)

	defer k.readline.Close()
	log.SetOutput(k.readline.Stderr())

	for k.run {
		line, err = k.readline.Readline()
		scope = ""

		if err == readline.ErrInterrupt || err == io.EOF {
			return
		}

		if len(line) == 0 {
			continue
		}

		if fkc != nil {
			k._cleanRequest(request)
			scope = fkc.GetScope()
			fkc = nil
		}

		kc = CreateCmd(line, scope)
		fkc, err = k.Exec(kc, request)

		if err != nil {
			fkc.WriteError(err)
			continue
		}

		fkc.WriteSuccess(fkc.GetResult())
	}
}

func (k *Kernel) _initReadline(url string) error {
	var (
		rl            *readline.Instance
		config        *readline.Config
		autocompleter *readline.PrefixCompleter
		commands      []string
		err           error
	)

	autocompleter = readline.NewPrefixCompleter()
	commands = append(k.commons, k.GetItemsName()...)

	for _, item := range commands {
		child := readline.PcItem(item)
		autocompleter.SetChildren(append(autocompleter.GetChildren(), child))
	}

	config = &readline.Config{
		Prompt:          "(" + url + ")$ ",
		HistoryFile:     "/tmp/razboynik.tmp",
		AutoComplete:    autocompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	rl, err = readline.NewEx(config)
	k.readline = rl

	return err
}

func _kernelDefault(kc *KernelCmd, request *core.REQUEST) (*KernelCmd, error) {
	return nil, errors.New("No default fonction defined")
}
