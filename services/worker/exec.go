package worker

import (
	"fmt"

	"github.com/eatbytes/razboy/core"
)

func Exec(cmd string, config *core.SERVERCONFIG) {
	fmt.Println(cmd)
}
