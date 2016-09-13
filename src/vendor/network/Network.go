package network

import (
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
	url     string
	method  string
	headers []string
	form    string
	jar     []string
	proxy   string
	status  bool
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
	return ""
}

func (n *NETWORK) _initConfig(r string) {
	c := config{
		url:    n.host,
		method: "GET",
		status: true,
	}

	c.headers = []string{"UserAgent"}
	request := r + n._getHandleBack()

	if n.method == 0 { //GET
		c.url = n.host + "?" + n.parameter + "=" + request
	} else if n.method == 1 { //POST
		c.form = request
		c.method = "POST"
	} else if n.method == 2 { //HEADER
		c.headers = []string{request}
	} else if n.method == 3 { //COOKIE
		c.jar = []string{}
	} else {
		c.status = false
	}

	n._config = c
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

	req, err := http.NewRequest(n._config.method, n._config.url, nil)

	if err != nil {
		fmt.Println("HERE 2")
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("HERE 3")
		panic(err)
	}

	n._handleSuccess(resp)
}
