package syscall

import (
	"fmt"

	"github.com/eatbytes/sysgo"
)

func Syscall(str string) {
	result, err := sysgo.Call(str)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
