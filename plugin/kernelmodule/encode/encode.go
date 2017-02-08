package main

import (
	"fmt"
	"strings"

	"os"

	"github.com/eatbytes/razboy"
)

func main() {
	str := strings.Join(os.Args[1:len(os.Args)], " ")
	sEnc := razboy.Encode(str)

	fmt.Println(sEnc)
}
