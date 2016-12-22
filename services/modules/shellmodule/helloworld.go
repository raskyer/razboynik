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
	Flags  *HelloWorldFlags
	Stdout string
	Stderr string
}

func (hw *HelloWorldCmd) Init(stdout, stderr string) {
	hw.Stdout = stdout
	hw.Stderr = stderr
	hw.Flags = new(HelloWorldFlags)
}

func (hw *HelloWorldCmd) Exec(kl *kernel.KernelLine, config *razboy.Config) (kernel.KernelCommand, int, error) {
	var (
		parser  *flags.Parser
		err, we error
	)

	parser = flags.NewParser(hw.Flags, flags.Default)
	parser.Name = hw.GetName()

	_, err = parser.ParseArgs(kl.GetArr())

	if err != nil {
		we = hw.WriteError(err)

		if we != nil {
			return nil, 1, we
		}

		return hw, 1, err
	}

	we = hw.WriteSuccess("Hello", hw.Flags.Name)

	return hw, 0, we
}

func (hw *HelloWorldCmd) Write(e error, i ...interface{}) error {
	if e != nil {
		return hw.WriteError(e)
	}

	return hw.WriteSuccess(i...)
}

func (hw *HelloWorldCmd) WriteSuccess(i ...interface{}) error {
	return kernel.Boot().WriteSuccess(hw.Stdout, i...)
}

func (hw *HelloWorldCmd) WriteError(e error) error {
	return kernel.Boot().WriteError(hw.Stderr, e)
}

func (hw *HelloWorldCmd) GetName() string {
	return "-helloworld"
}

func (hw *HelloWorldCmd) GetCompleter() (kernel.CompleteFunction, bool) {
	return hw.Complete, true
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
