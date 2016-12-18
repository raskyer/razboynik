package kernel

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/eatbytes/razboy"
)

func (k *Kernel) ListLocalFiles(line string, c *razboy.Config) []string {
	return Local(line)
}

func Local(line string, plus ...string) []string {
	var (
		lenPlus int
		arr     []string
		names   []string
		tmp     []string
		dir     string
	)

	lenPlus = len(plus)
	names = make([]string, 0)
	dir = ""

	if lenPlus < 1 {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
		arr = strings.Fields(line)

		if len(arr) > 1 {
			dir = arr[1]
		}
	} else {
		arr = append(plus[:0], plus[:len(plus)-1]...)
		dir = strings.Join(arr, "/")
	}

	files, _ := ioutil.ReadDir(dir)

	if len(arr) > 1 {
		arr = strings.Split(arr[1], "/")
		tmp = Local(line, arr...)

		if len(tmp) > 0 {
			return tmp
		}
	}

	for _, f := range files {
		name := f.Name()

		if lenPlus > 0 {
			if !strings.Contains(name, plus[lenPlus-1]) {
				continue
			}

			name = strings.Replace(name, plus[lenPlus-1], "", -1)
		}

		names = append(names, name)
	}

	return names
}

func (k *Kernel) ListRemoteFiles(line string, c *razboy.Config) []string {
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
