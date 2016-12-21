package targetwork

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/usr"
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
	fmt.Print("Enter target's " + color.YellowString(txt) + " (\"" + color.MagentaString(def) + "\"): ")

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

	return tmp
}

func _getInputBool(txt string, def bool) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target's " + color.YellowString(txt) + " (\"" + color.MagentaString(strconv.FormatBool(def)) + "\"): ")

	tmp, s, err := reader.ReadRune()

	if err != nil || s > 1 {
		fmt.Println("Please answer : y or n")
		fmt.Println(err)
		return _getInputBool(txt, def)
	}

	if tmp == 'y' {
		color.Green("true")
		return true
	} else if tmp == 'n' || tmp == '\n' {
		color.Green("false")
		return false
	} else {
		fmt.Println("Please answer : y or n")
		return _getInputBool(txt, def)
	}
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
	target.Config.Method = _getInput("method ['GET', 'POST', 'HEADER', 'COOKIE']", target.Config.Method)
	target.Config.Parameter = _getInput("parameter", target.Config.Parameter)
	target.Config.Shellmethod = _getInput("shell method ['system', 'shell_exec', 'proc_open', 'passthru']", target.Config.Shellmethod)
	target.Config.Shellscope = _getInput("shell scope ['./', '/']", target.Config.Shellscope)
	target.Config.Key = _getInput("key", target.Config.Key)
	target.Config.NoExtra = _getInputBool("no-extra ['y', 'n']", target.Config.NoExtra)
}

func FindTarget(config *Configuration, name string) (*Target, int, error) {
	for i, item := range config.Targets {
		if item.Name == name {
			return item, i, nil
		}
	}

	return nil, 0, errors.New("No available target with name: " + name)
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
		dir string
		err error
	)

	dir, err = usr.GetHomeDir()

	if err != nil {
		return "", err
	}

	return dir + "/.razboynik.json", nil
}

func CreateFile(filepath string) error {
	var (
		file *os.File
		err  error
	)

	file, err = os.Create(filepath)
	defer file.Close()

	if err != nil {
		return err
	}

	err = SaveConfiguration(new(Configuration))

	return err
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

	if os.IsNotExist(err) {
		err = CreateFile(filepath)
	}

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
