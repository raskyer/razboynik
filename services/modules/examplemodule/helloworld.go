package examplemodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	pflag "github.com/spf13/pflag"
)

type HelloWorldCmd struct {
	Name string
}

func (hw *HelloWorldCmd) InitFlags(args []string) {
	flaghandler := pflag.NewFlagSet("helloworld", pflag.ContinueOnError)
	flaghandler.StringVarP(&hw.Name, "name", "n", "world", "Name your hello")
	flaghandler.Parse(args)
}

func (hw *HelloWorldCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	hw.InitFlags(kl.GetArr())
	kl.WriteSuccess("Hello", hw.Name)

	return hw, nil
}

func (hw *HelloWorldCmd) GetName() string {
	return "-helloworld"
}

func (hw *HelloWorldCmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return hw.Complete, false
}

func (hw *HelloWorldCmd) Complete(l string, c *razboy.Config) []string {
	return []string{"--name"}
}

func (hw *HelloWorldCmd) GetResult() []byte {
	return make([]byte, 0)
}

func (hw *HelloWorldCmd) GetResultStr() string {
	return ""
}
