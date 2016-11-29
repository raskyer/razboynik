package targetwork

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/eatbytes/razboy"
)

type Target struct {
	Name   string         `json:"name"`
	Config *razboy.Config `json:"config"`
}

type Configuration struct {
	Targets []*Target `json:"targets"`
}

func _getInput(txt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter '" + txt + "' of target: ")
	tmp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return _getInput(txt)
	}

	return tmp
}

func CreateTarget() *Target {
	var (
		target *Target
	)

	target = new(Target)
	target.Name = _getInput("Name")
	target.Config = new(razboy.Config)
	target.Config.Url = _getInput("URL")
	target.Config.Method = _getInput("Method [GET, POST, HEADER, COOKIE]")

	fmt.Println(target)

	return target
}

func SaveConfiguration(config *Configuration) error {
	var (
		filepath string
		buffer   []byte
		err      error
	)

	buffer, err = json.Marshal(config)

	if err != nil {
		return err
	}

	filepath, err = GetConfigFilePath()

	if err != nil {
		return err
	}

	return nil

	return ioutil.WriteFile(filepath, buffer, 0644)
}

func GetConfigFilePath() (string, error) {
	var (
		usr *user.User
		err error
	)

	usr, err = user.Current()

	if err != nil {
		return "", err
	}

	return usr.HomeDir + "/.razboynik.json", nil
}

func GetConfiguration() (*Configuration, error) {
	var (
		file     *os.File
		decoder  *json.Decoder
		config   *Configuration
		filepath string
		err      error
	)

	filepath, err = GetConfigFilePath()

	if err != nil {
		return config, err
	}

	config = new(Configuration)
	file, err = os.Open(filepath)

	if err != nil {
		return config, err
	}

	decoder = json.NewDecoder(file)
	err = decoder.Decode(config)

	if err != nil {
		return config, err
	}

	return config, nil
}
