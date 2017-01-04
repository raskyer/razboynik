package main

import (
	"log"

	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/spf13/pflag"
)

type Cmd struct{}

var name string

func main() {
	provider, err := kernel.CreateProvider(Cmd{})

	if err != nil {
		log.Fatalf("Failed to register Plugin: %s", err)
	}

	provider.Serve()
}

func InitFlags(args []string) {
	flaghandler := pflag.NewFlagSet("helloworld", pflag.ContinueOnError)
	flaghandler.StringVarP(&name, "name", "n", "world", "Name your hello")
	flaghandler.Parse(args)
}

func (Cmd) Exec(args kernel.KernelExternalArgs, response *string) error {
	kl := kernel.CreateLine(args.Line)
	InitFlags(kl.GetArg())

	*response = "Hello " + name

	return nil
}
