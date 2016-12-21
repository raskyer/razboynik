package kernel

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/eatbytes/razboy"
)

func (k *Kernel) ListLocalFiles(line string, c *razboy.Config) []string {
	return Local(line, "")
}

func Local(line string, filename string) []string {
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
	arg = strings.TrimPrefix(line, a[0]+" ")
	arg = strings.TrimSpace(arg)

	f, err := os.Stat(arg)
	if err == nil {
		if f.Mode().IsRegular() {
			return []string{}
		}

		if f.IsDir() && !strings.HasSuffix(arg, "/") {
			return []string{"/"}
		}
	}

	if arg != "" {
		dir = arg
	} else {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	items, err = ioutil.ReadDir(dir)

	//Check if former directory as other result
	if err != nil && filename == "" {
		arrScope = strings.Split(arg, "/")
		filename = arrScope[len(arrScope)-1]
		scope = strings.TrimSuffix(arg, filename)

		tmp := Local(scope, filename)

		if len(tmp) > 0 {
			return tmp
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

func (k *Kernel) ListRemoteFiles(line string, c *razboy.Config) []string {
	if c.NoExtra {
		return make([]string, 0)
	}

	return k.Remote(line, c)
}

func (k *Kernel) Remote(line string, c *razboy.Config) []string {
	var addScope string

	arr := strings.Fields(line)

	if len(arr) > 1 {
		addScope = arr[1]
	}

	kc := CreateCmd("ls " + addScope)
	kc, err := k.Exec(kc, c)

	if err != nil {
		return make([]string, 0)
	}

	f := kc.GetResult()

	return strings.Fields(f)
}

func (k *Kernel) ListRemoteFilesPHP(line string, c *razboy.Config) []string {
	if c.NoExtra {
		return make([]string, 0)
	}

	return k.RemotePHP(line, c)
}

func (k *Kernel) RemotePHP(line string, c *razboy.Config) []string {
	var addScope string

	arr := strings.Fields(line)

	if len(arr) > 1 {
		addScope = arr[1]
	}

	kc := CreateCmd("-listfile " + addScope)
	kc, err := k.Exec(kc, c)

	if err != nil {
		return make([]string, 0)
	}

	f := kc.GetResult()

	return strings.Fields(f)
}
