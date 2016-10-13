package bash

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()
var keeper string

func init() {
	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	clear["windows"] = func() {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	fmt.Printf("\n\n\n\n\n")

	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}

	if keeper != "" {
		fmt.Println(keeper)
	}
}

func SetKeeper(str string) {
	keeper = str
}
