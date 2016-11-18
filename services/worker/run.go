package worker

import (
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Run(config *core.REQUEST) error {
	var (
		rl  *readline.Instance
		err error
	)

	rl, err = _run_getReadline(config.SRVc.Url)

	if err != nil {
		return err
	}

	defer rl.Close()
	log.SetOutput(rl.Stderr())

	_run_loop(rl, config)

	return nil
}

func _run_getReadline(url string) (*readline.Instance, error) {
	var (
		k             *kernel.Kernel
		config        *readline.Config
		autocompleter *readline.PrefixCompleter
		commands      []string
	)

	k = kernel.Boot()

	autocompleter = readline.NewPrefixCompleter()
	commands = append(k.GetCommons(), k.GetItemsName()...)

	for _, item := range commands {
		child := readline.PcItem(item)
		autocompleter.SetChildren(append(autocompleter.GetChildren(), child))
	}

	config = &readline.Config{
		Prompt:          "(" + url + ")$ ",
		HistoryFile:     "/tmp/readlinebash.tmp",
		AutoComplete:    autocompleter,
		InterruptPrompt: "^C",
		EOFPrompt:       "exit",
	}

	return readline.NewEx(config)
}

func _run_cleanRequest(config *core.REQUEST) {
	config.Type = ""
	config.Raw = ""
}

func _run_loop(rl *readline.Instance, config *core.REQUEST) error {
	for true {
		line, err := rl.Readline()

		if err == readline.ErrInterrupt || err == io.EOF || line == "exit" {
			return nil
		}

		if len(line) == 0 {
			continue
		}

		_run_cleanRequest(config)
		config.SHLc.Cmd = line

		kc, err := Exec(config)

		if err != nil {
			kc.WriteError(err)
			continue
		}

		if kc.GetScope() != config.SHLc.Scope {
			config.SHLc.Scope = kc.GetScope()
			rl.SetPrompt("(" + config.SRVc.Url + "):" + kc.GetScope() + "$ ")
		}

		kc.WriteSuccess(kc.GetResult())
	}

	return nil
}
