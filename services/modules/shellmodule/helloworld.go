package shellmodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/jessevdk/go-flags"
)

type HelloWorldFlags struct {
	Name string `short:"n" long:"name" description:"Name"`
}

type HelloWorldCmd struct {
	Flags *HelloWorldFlags
}

func (hw *HelloWorldCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, error) {
	var (
		parser  *flags.Parser
		err, we error
	)

	parser = flags.NewParser(hw.Flags, flags.Default)
	parser.Name = hw.GetName()

	_, err = parser.ParseArgs(kl.GetArr())

	if err != nil {
		we = kl.WriteError(err)

		if we != nil {
			return hw, we
		}

		return hw, err
	}

	we = kl.WriteSuccess("Hello", hw.Flags.Name)

	return hw, we
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
