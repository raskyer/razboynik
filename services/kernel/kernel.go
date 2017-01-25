package kernel

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"reflect"
	"strings"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboy"
)

type CompleterFunction func(string, *razboy.Config) []string
type GetCompleterFunction func() (CompleterFunction, bool)

type Kernel struct {
	def           *Item
	commands      []*Item
	readline      *readline.Instance
	autocompleter *readline.PrefixCompleter
	run           bool
}

var kInstance *Kernel
var kSilent bool

func Boot() *Kernel {
	if kInstance == nil {
		kInstance = new(Kernel)
	}

	return kInstance
}

func Silent() {
	kSilent = true
}

func (k *Kernel) Exec(line string, config *razboy.Config) Response {
	return k.ExecKernelLine(CreateLine(line), config)
}

func (k *Kernel) ExecKernelLine(l *Line, config *razboy.Config) Response {
	for _, cmd := range k.commands {
		if cmd.Name == l.GetName() {
			return cmd.Exec(l, config)
		}
	}

	if k.def != nil {
		return k.def.Exec(l, config)
	}

	return k.Default(l, config)
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
		r    Response
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

		r = k.Exec(line, config)

		if r.Err != nil {
			fmt.Println(r.Err.Error())
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
		if item.Completer == nil {
			child = readline.PcItem(item.Name)
			children = append(children, child)
			continue
		}

		completer, multilevel := item.Completer()

		child = readline.PcItem(
			item.Name,
			readline.PcItemDynamic(dynamicAdapter(completer, c)),
		)

		if multilevel {
			child.MultiLevel = true
		}

		children = append(children, child)
	}

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

func (k Kernel) Default(l *Line, config *razboy.Config) Response {
	return Response{Err: errors.New("No default fonction defined")}
}

func (k Kernel) GetCommands() []*Item {
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

func (k *Kernel) SetDefault(d *Item) {
	k.def = d
}

func (k *Kernel) SetCommands(items []*Item) {
	k.commands = items
}

func Write(stdout, stderr *os.File, e error, i ...interface{}) error {
	if kSilent {
		return nil
	}

	if e != nil {
		return WriteError(stderr, e)
	}

	return WriteSuccess(stdout, i...)
}

func WriteSuccess(stdout *os.File, i ...interface{}) error {
	var e error

	if kSilent {
		return nil
	}

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
	if kSilent {
		return nil
	}

	_, e := fmt.Fprintln(stderr, err.Error())

	return e
}

func dynamicAdapter(completer CompleterFunction, c *razboy.Config) func(string) []string {
	return func(line string) []string {
		return completer(line, c)
	}
}
