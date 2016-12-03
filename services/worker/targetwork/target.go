package targetwork

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"strings"

	"errors"

	"github.com/eatbytes/razboy"
	"github.com/fatih/color"
)

type Target struct {
	Name   string         `json:"name"`
	Config *razboy.Config `json:"config"`
}

type Configuration struct {
	Targets []*Target `json:"targets"`
}

func _getInput(txt, def string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target's " + color.GreenString(txt) + " (\"" + color.MagentaString(def) + "\"): ")

	tmp, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return _getInput(txt, def)
	}

	tmp = strings.TrimSpace(tmp)

	if tmp == "" {
		tmp = def
	} else if tmp == "x" {
		tmp = ""
	}

	color.Green(tmp)

	return def
}

func CreateTarget() *Target {
	var (
		target *Target
	)

	target = new(Target)
	target.Config = razboy.NewConfig()
	EditTarget(target)

	return target
}

func EditTarget(target *Target) {
	target.Name = _getInput("name", target.Name)
	target.Config.Url = _getInput("URL", target.Config.Url)
	target.Config.Method = _getInput("method [GET, POST, HEADER, COOKIE]", target.Config.Method)
	target.Config.Parameter = _getInput("parameter", target.Config.Parameter)
	target.Config.Shellmethod = _getInput("shell method [system, shell_exec]", target.Config.Shellmethod)
	target.Config.Key = _getInput("key", target.Config.Key)
}

func FindTarget(config *Configuration, name string) (*Target, error) {
	for _, item := range config.Targets {
		if item.Name == name {
			return item, nil
		}
	}

	return nil, errors.New("No available target with name: " + name)
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

	defer file.Close()

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
