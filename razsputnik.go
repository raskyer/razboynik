package razsputnik

import (
	"github.com/eatbytes/razboy"
	"github.com/smallnest/rpcx"
)

const (
	PROTOCOL = "tcp"
	ADDR     = ":8972"
	OBJECT   = "Kernel"
)

type Args interface{}

func buildClient() *rpcx.Client {
	selector := &rpcx.DirectClientSelector{Network: PROTOCOL, Address: ADDR}
	return rpcx.NewClient(selector)
}

func GetConfig() (*razboy.Config, error) {
	config := new(razboy.Config)
	client := buildClient()
	err := client.Call(OBJECT+".GetConfig", nil, config)

	return config, err
}

func UpdateConfig(config *razboy.Config) error {
	client := buildClient()

	return client.Call(OBJECT+".UpdateConfig", config, nil)
}

func UpdatePrompt(str string) error {
	client := buildClient()
	return client.Call(OBJECT+".UpdatePrompt", &str, nil)
}
