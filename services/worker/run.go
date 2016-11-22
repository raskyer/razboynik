package worker

import "github.com/eatbytes/razboy/core"

func Run(request *core.REQUEST) error {
	var (
		b   *Bash
		err error
	)

	b, err = CreateBash(request)

	if err != nil {
		return err
	}

	b.loop()

	return nil
}
