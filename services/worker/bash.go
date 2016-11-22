package worker

import (
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/kernel"
)

type Bash struct {
	request  *core.REQUEST
	readline *readline.Instance
	run      bool
}

func CreateBash(request *core.REQUEST) (*Bash, error) {
	var (
		rl  *readline.Instance
		err error
	)

	rl, err = createReadline(request.SRVc.Url)

	if err != nil {
		return nil, err
	}

	return &Bash{
		request:  request,
		readline: rl,
		run:      true,
	}, nil
}

func createReadline(url string) (*readline.Instance, error) {
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

func (b *Bash) cleanRequest() {
	b.request.Type = ""
	b.request.Raw = ""
}

func (b *Bash) loop() {
	defer b.readline.Close()
	log.SetOutput(b.readline.Stderr())

	for b.run {
		line, err := b.readline.Readline()

		if err == readline.ErrInterrupt || err == io.EOF {
			return
		}

		if len(line) == 0 {
			continue
		}

		b.cleanRequest()

		command, err := Exec(line, b.request)

		if err != nil {
			command.WriteError(err)
			continue
		}

		command.WriteSuccess(command.GetResult())
	}
}

func (b *Bash) UpdatePrompt(scope string) {
	b.readline.SetPrompt("(" + b.request.SRVc.Url + "):" + scope + "$ ")
}
