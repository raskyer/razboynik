package bash

import (
	"fmt"
	"strings"
)

type BashCommand struct {
	raw string
	arr []string
	str string
	out string
	err string
	Fn  spFunc
}

func (bc *BashCommand) Write(str string, err error) {
	if err != nil {
		bc.WriteError(err)
		return
	}

	bc.WriteSuccess(str)
}

func (bc *BashCommand) WriteSuccess(str string) {
	if bc.out == "1" {
		fmt.Println(str)
	}
}

func (bc *BashCommand) WriteError(err error) {
	if bc.err == "2" {
		fmt.Println(err.Error())
	}
}

func (b *BashInterface) CreateCommand(raw string) *BashCommand {
	arr := strings.Fields(raw)

	fnInt := findFunc(arr[0], b.specialCmd)
	fn := b.specialFunc[fnInt]

	strArr := append(arr[1:], arr[len(arr):]...)
	str := strings.Join(strArr, " ")

	out := findOut(raw, arr)
	err := findErr(raw, arr)

	cmd := BashCommand{
		raw: raw,
		arr: arr,
		str: str,
		out: out,
		err: err,
		Fn:  fn,
	}

	return &cmd
}

func findOut(str string, arr []string) string {
	if strings.Contains(str, ">") {

	}

	return "1"
}

func findErr(str string, arr []string) string {
	if strings.Contains(str, "2>") {

	}

	return "2"
}

func findFunc(str string, cmds []string) int {
	for i, item := range cmds {
		if str == item {
			return i
		}
	}

	return 0
}
