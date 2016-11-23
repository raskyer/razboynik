package kernel

import (
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy/core"
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
		k             *Kernel
		config        *readline.Config
		autocompleter *readline.PrefixCompleter
		commands      []string
	)

	k = Boot()

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
	b.request.Action = ""
}

func (b *Bash) loop() {
	var (
		kc   *KernelCmd
		line string
		err  error
	)

	defer b.readline.Close()
	log.SetOutput(b.readline.Stderr())

	for b.run {
		line, err = b.readline.Readline()

		if err == readline.ErrInterrupt || err == io.EOF {
			return
		}

		if len(line) == 0 {
			continue
		}

		b.cleanRequest()

		kc = CreateCmd(line, b.request.SHLc.Scope)
		kc, err = kc.Exec(b.request)

		if err != nil {
			kc.WriteError(err)
			continue
		}

		kc.WriteSuccess(kc.GetResult())
	}
}

func (b *Bash) Run() {
	b.loop()
}

func (b *Bash) UpdatePrompt(scope string) {
	b.readline.SetPrompt("(" + b.request.SRVc.Url + "):" + scope + "$ ")
}
