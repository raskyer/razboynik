package configuration

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"github.com/eatbytes/razboynik/pkg/services/worker/target"
	"github.com/eatbytes/razboynik/pkg/services/worker/usr"
)

type Configuration struct {
	Targets   []*target.Target `json:"targets"`
	PluginDir string           `json:"plugin_dir"`
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

func FindTarget(config *Configuration, name string) (*target.Target, int, error) {
	for i, item := range config.Targets {
		if item.Name == name {
			return item, i, nil
		}
	}

	return nil, 0, errors.New("No available target with name: " + name)
}
