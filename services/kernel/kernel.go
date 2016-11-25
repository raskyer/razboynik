package kernel

import (
	"io"
	"log"

	"github.com/chzyer/readline"
	"github.com/eatbytes/razboynik/services/config"
)

type KernelFunction func(*KernelCmd, *config.Config) (*KernelCmd, error)

type KernelItem struct {
	Name string
	Fn   KernelFunction
}

type Kernel struct {
	def      *KernelItem
	items    []*KernelItem
	readline *readline.Instance
	former   *KernelCmd
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

func (k *Kernel) Exec(kc *KernelCmd, config *config.Config) (*KernelCmd, error) {
	for _, item := range k.items {
		if item.Name == kc.name {
			return item.Fn(kc, config)
		}
	}

	return k.def.Fn(kc, config)
}

func (k *Kernel) Run(config *config.Config) error {
	var err error

	err = k._initReadline(config.Url)

	if err != nil {
		return err
	}

	k.run = true
	k._loop(config)

	return nil
}

func (k *Kernel) _loop(config *config.Config) {
	var (
		kc, fkc     *KernelCmd
		line, scope string
		err         error
	)

	scope = ""

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
			scope = fkc.GetScope()
			k.SetFormerCmd(fkc)
		}

		kc = CreateCmd(line, scope)
		fkc, err = k.Exec(kc, config)

		if err != nil {
			fkc.WriteError(err)
			continue
		}

		fkc.WriteSuccess(fkc.GetResult())
	}
}
