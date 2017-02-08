package usr

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

func GetHomeDir() (string, error) {
	var (
		usr *user.User
		err error
	)

	usr, err = user.Current()

	if err != nil {
		return "", err
	}

	return usr.HomeDir, nil
}

func ListDir(dir string) []string {
	var (
		list  []string
		items []os.FileInfo
		err   error
	)

	if dir == "" {
		dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	}

	items, err = ioutil.ReadDir(dir)

	if err != nil {
		return list
	}

	for _, item := range items {
		list = append(list, item.Name())
	}

	return list
}
