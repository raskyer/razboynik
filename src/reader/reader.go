package reader

import (
	"fmt"
	"fuzzer"
)

func ReadEncode(str string) {
	sDec := fuzzer.Decode(str)
	fmt.Println(sDec)
}

func Read(str string) {
	fmt.Println(str)
}
