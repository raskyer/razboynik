package shellmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type HelloWorldCmd struct{}

func (hw *HelloWorldCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	kl.WriteSuccess("Hello", "World")

	return hw, nil
}

func (hw *HelloWorldCmd) GetName() string {
	return "-helloworld"
}

func (hw *HelloWorldCmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return hw.Complete, false
}

func (hw *HelloWorldCmd) Complete(l string, c *razboy.Config) []string {
	return []string{"\"John doe\"", "\"Marry Jane\""}
}

func (hw *HelloWorldCmd) GetResult() []byte {
	return make([]byte, 0)
}

func (hw *HelloWorldCmd) GetResultStr() string {
	return ""
}
