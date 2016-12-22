package kernel

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func (k Kernel) GetCommands() []KernelCommand {
	return k.commands
}

func (k Kernel) Write(stdout, stderr string, e error, i ...interface{}) error {
	if e != nil {
		return k.WriteError(stderr, e)
	}

	return k.WriteSuccess(stdout, i)
}

func (k Kernel) WriteSuccess(stdout string, i ...interface{}) error {
	var isString = true

	for _, v := range i {
		if reflect.TypeOf(v).Kind() == reflect.String {
			isString = true
		}

		if stdout != "&1" && stdout != "" {
			if isString {
				return k.WriteInFile(stdout, []byte(reflect.TypeOf(v).String()))
			}
		} else {
			if isString {
				fmt.Print(strings.TrimSpace(v.(string)), " ")
			}
		}
	}

	fmt.Print("\n")

	return nil
}

func (k Kernel) WriteError(stderr string, err error) error {
	if stderr != "&2" {
		return k.WriteInFile(stderr, []byte(err.Error()))
	}

	fmt.Println(err.Error())

	return nil
}

func (k Kernel) WriteInFile(path string, buf []byte) error {
	var (
		f   *os.File
		err error
	)

	f, err = os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(buf)

	return err
}

func (k *Kernel) StartRun() {
	k.run = true
}

func (k *Kernel) StopRun() {
	k.run = false
}

func (k *Kernel) UpdatePrompt(url, scope string) {
	if k.readline == nil {
		return
	}

	k.readline.SetPrompt("(" + url + "):" + scope + "$ ")
}

func (k *Kernel) SetDefault(d KernelCommand) {
	k.def = d
}

func (k *Kernel) SetCommands(cmd []KernelCommand) {
	k.commands = cmd
}
