package kernelmodule

import (
	"os"
	"strings"

	"fmt"

	"github.com/eatbytes/razboy"
)

func main() {
	str := strings.Join(os.Args[1:len(os.Args)], " ")
	sDec, err := razboy.Decode(str)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	fmt.Println(sDec)
}
