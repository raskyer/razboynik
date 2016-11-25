package kernel

import (
	"errors"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboynik/services/config"
)

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

func _kernelDefault(kc *KernelCmd, config *config.Config) (*KernelCmd, error) {
	return kc, errors.New("No default fonction defined")
}
