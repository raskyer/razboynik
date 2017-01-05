package main

import (
	"log"

	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/provider"
)

func main() {
	provider, err := provider.CreateProvider(PLUG{})

	if err != nil {
		log.Fatalf("Failed to register Plugin: %s", err)
	}

	provider.Serve()
}

type PLUG struct{}

func (PLUG) Exec(args *provider.Args, resp *provider.Response) error {
	kl := kernel.CreateLine(args.Line)
	name := "world"
	if len(kl.GetArg()) > 0 {
		name = kl.GetArg()[0]
	}

	resp.Code = provider.PRINT_CODE
	resp.Content = "Hello " + name

	return nil
}

func (PLUG) Completer(args *provider.Args, resp *provider.Response) error {
	resp.Items = []string{"World", "John Doe"}

	return nil
}
