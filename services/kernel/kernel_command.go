package kernel

import "github.com/eatbytes/razboy"

type KernelCommand interface {
	Exec(*KernelLine, *razboy.Config) KernelResponse
	GetName() string
	GetCompleter() (CompleterFunction, bool)
}
