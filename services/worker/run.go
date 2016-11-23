package worker

import (
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/bash"
)

func Run(request *core.REQUEST) error {
	var (
		b   *bash.Bash
		err error
	)

	b, err = bash.CreateBash(request)

	if err != nil {
		return err
	}

	b.Run()

	return nil
}
