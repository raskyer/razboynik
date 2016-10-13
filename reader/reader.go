package reader

import (
	"fmt"

	"github.com/eatbytes/fuzzcore"
)

func ReadOne(r, msg string) {
	if r == "1" {
		fmt.Println(msg)
		return
	}

	fmt.Println("An error occured")
}

func ReadEncode(str string) {
	sDec := fuzzcore.Decode(str)
	fmt.Println(sDec)
}

func Read(str string) {
	fmt.Println(str)
}
