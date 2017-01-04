package examplemodule

import (
	"os"

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

func (hw *HelloWorldCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	hw.InitFlags(kl.GetArr())
	kernel.WriteSuccess(os.Stdout, "Hello", hw.Name)

	return nil
}

func (hw *HelloWorldCmd) GetName() string {
	return "-helloworld"
}

func (hw *HelloWorldCmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return hw.Complete, false
}

func (hw *HelloWorldCmd) Complete(l string, c *razboy.Config) []string {
	return []string{"--name"}
}
