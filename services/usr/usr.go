package usr

import "os/user"

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
