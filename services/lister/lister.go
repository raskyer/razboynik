package lister

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

func Local(line string, config *razboy.Config) []string {
	return local(line, "")
}

func local(line string, filename string) []string {
	var (
		arrScope []string
		names    []string
		dir      string
		arg      string
		scope    string
		items    []os.FileInfo
		err      error
	)

	a := strings.Fields(line)

	if len(a) < 1 {
		return []string{}
	}

	arg = strings.TrimPrefix(line, a[0]+" ")
	arg = strings.TrimSpace(arg)

	//Generate a bug, TO DO: fix this better
	if strings.Contains(arg, "~") {
		return []string{}
	}

	//Check if path is a valid folder or file
	f, err := os.Stat(arg)
	if err == nil {
		if f.Mode().IsRegular() {
			return []string{}
		}

		if f.IsDir() && !strings.HasSuffix(arg, "/") {
			return []string{"/"}
		}
	}

	//Choose the right path
	dir = arg
	if dir == "" {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	//Get items
	items, err = ioutil.ReadDir(dir)

	//Check if former directory as other result
	if err != nil && filename == "" {
		arrScope = strings.Split(arg, "/")

		if len(arrScope) > 0 {
			filename = arrScope[len(arrScope)-1]
			scope = strings.TrimSuffix(arg, filename)

			formerDirectory := local(scope, filename)
			if len(formerDirectory) > 0 {
				return formerDirectory
			}
		}
	}

	names = []string{}
	for _, i := range items {
		name := i.Name()

		if filename != "" {
			if !strings.HasPrefix(name, filename) {
				continue
			}

			name = strings.Replace(name, filename, "", -1)
		}

		if name == "" {
			continue
		}

		f, err := os.Stat(dir + "/" + filename + name)
		if err == nil && f.IsDir() {
			name = name + "/"
		} else {
			name = name + " "
		}

		names = append(names, name)
	}

	if len(names) > 0 {
		names = append(names, "../", "./")
	}

	return names
}

func RemoteSHELL(line string, config *razboy.Config) []string {
	var (
		addScope string
		arr      []string
		err      error
		cmd      kernel.KernelCommand
	)

	arr = strings.Fields(line)

	if len(arr) > 1 {
		addScope = arr[1]
	}

	cmd, err = kernel.Boot().Exec("ls "+addScope+" -> /dev/null", config)

	if err != nil {
		return make([]string, 0)
	}

	return strings.Fields(cmd.GetResultStr())
}

func RemotePHP(line string, config *razboy.Config) []string {
	var (
		addScope string
		arr      []string
		err      error
		cmd      kernel.KernelCommand
	)

	arr = strings.Fields(line)

	if len(arr) > 1 {
		addScope = arr[1]
	}

	cmd, err = kernel.Boot().Exec("-list "+addScope+" -> /dev/null", config)

	if err != nil {
		return make([]string, 0)
	}

	return strings.Fields(cmd.GetResultStr())
}
