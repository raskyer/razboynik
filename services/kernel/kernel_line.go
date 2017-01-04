package kernel

import (
	"fmt"
	"os"
	"strings"

	shellwords "github.com/mattn/go-shellwords"
)

type KernelLine struct {
	name string
	raw  string
	str  string
	arg  []string
	out  *os.File
	err  *os.File
}

func CreateLine(raw string) *KernelLine {
	var (
		arg       []string
		name, str string
		e         error
		out, err  *os.File
	)

	out = os.Stdout
	err = os.Stderr

	arg, e = shellwords.Parse(raw)

	if e != nil {
		arg = strings.Fields(raw)
	}

	if len(arg) > 0 {
		name = arg[0]
		arg = append(arg[1:], arg[len(arg):]...)
	}

	if i := extractIn(arg, "->"); i != -1 {
		out, e = os.OpenFile(arg[i+1], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)

		if e != nil {
			fmt.Println(e)
			out = os.Stdout
		}

		arg = append(arg[:i], arg[i+2:]...)
		raw = name + " " + strings.Join(arg, " ")
	}

	if i := extractIn(arg, "-2>"); i != -1 {
		err, e = os.OpenFile(arg[i+1], os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)

		if e != nil {
			fmt.Println(e)
			err = os.Stderr
		}

		arg = append(arg[:i], arg[i+2:]...)
		raw = name + " " + strings.Join(arg, " ")
	}

	str = strings.Join(arg, " ")

	return &KernelLine{
		name: name,
		raw:  raw,
		str:  str,
		arg:  arg,
		out:  out,
		err:  err,
	}
}

func (kl KernelLine) GetRaw() string {
	return kl.raw
}

func (kl KernelLine) GetArg() []string {
	return kl.arg
}

func (kl KernelLine) GetArr() []string {
	return kl.arg
}

func (kl KernelLine) GetName() string {
	return kl.name
}

func (kl KernelLine) GetStr() string {
	return kl.str
}

func (kl KernelLine) GetStdout() *os.File {
	return kl.out
}

func (kl KernelLine) GetStderr() *os.File {
	return kl.err
}

func extractIn(slice []string, value string) int {
	for p, v := range slice {
		if v == value {
			return p
		}
	}

	return -1
}
