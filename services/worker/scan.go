package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Scan(config *razboy.Config) (*kernel.KernelCmd, error) {
	return Exec("-scan", config)
}
