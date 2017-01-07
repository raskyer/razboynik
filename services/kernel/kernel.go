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
	"github.com/eatbytes/razboynik/services/gflags"
	"github.com/eatbytes/razboynik/services/provider"
)

type CompleterFunction func(string, *razboy.Config) []string

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

func (k *Kernel) Exec(line string, config *razboy.Config) KernelResponse {
	return k.ExecKernelLine(CreateLine(line), config)
}

func (k *Kernel) ExecKernelLine(kl *KernelLine, config *razboy.Config) KernelResponse {
	if strings.HasPrefix(kl.GetName(), "#") {
		return k.ExecProvider(kl)
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

func (k *Kernel) ExecProvider(kl *KernelLine) KernelResponse {
	var (
		info *provider.Info
		args *provider.Args
		resp *provider.Response
		err  error
	)

	args = new(provider.Args)
	args.Line = kl.GetRaw()

	info = &provider.Info{
		Path:   provider.DIR + "/",
		Name:   strings.TrimPrefix(kl.GetName(), "#"),
		Method: provider.EXEC_FN,
	}

	resp, err = provider.CallProvider(info, args)

	//DEBUG
	if err != nil {
		WriteError(kl.GetStderr(), err)
	} else {
		WriteSuccess(kl.GetStdout(), resp.Content)
	}

	return KernelResponse{
		Err:  err,
		Body: []byte(resp.Content),
	}
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
		kr   KernelResponse
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

		kr = k.Exec(line, config)

		if kr.Err != nil {
			fmt.Println(kr.Err.Error())
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
	dir = dir + "/" + provider.DIR + "/"
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

func (k Kernel) Default(kl *KernelLine, config *razboy.Config) KernelResponse {
	return KernelResponse{Err: errors.New("No default fonction defined")}
}

func (k Kernel) GetCommands() []KernelCommand {
	return k.commands
}

func (k *Kernel) StartRun() {
	if gflags.Rpc {
		go LaunchRPC(&RPCServer{Port: 1234})
	}

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
		args := new(provider.Args)
		args.Line = line

		info := &provider.Info{
			Path:   path,
			Method: provider.COMPLETER_FN,
		}

		resp, err := provider.CallProvider(info, args)

		if err != nil {
			fmt.Println(err)
			return []string{}
		}

		return resp.Items
	}
}
