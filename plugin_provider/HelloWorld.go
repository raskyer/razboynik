package main

import (
	"log"

	"github.com/eatbytes/razboynik/services/kernel"
)

func main() {
	provider, err := kernel.CreateProvider(PLUG{})

	if err != nil {
		log.Fatalf("Failed to register Plugin: %s", err)
	}

	provider.Serve()
}

type PLUG struct{}

func (PLUG) Exec(args *kernel.KernelExternalArgs, resp *kernel.KernelExternalResponse) error {
	kl := kernel.CreateLine(args.Line)
	name := "world"

	if len(kl.GetArg()) > 0 {
		name = kl.GetArg()[0]
	}

	resp.Response = "Hello " + name

	return nil
}

func (PLUG) Completer(args *kernel.KernelExternalArgs, resp *kernel.KernelExternalResponse) error {
	resp.Items = []string{"World", "John Doe"}

	return nil
}
