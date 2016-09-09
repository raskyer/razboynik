package network

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli"
)

var NET = NETWORK{
	_host:      "http://localhost",
	_method:    0,
	_parameter: "fuzzer",
	_crypt:     false,
	_status:    false,
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
	_host       string
	_method     int
	_parameter  string
	_crypt      bool
	_httpMethod string
	_status     bool

	_config config
}

func (n *NETWORK) IsSetup() bool {
	return n._status
}

func (n *NETWORK) Setup(c *cli.Context) {
	url := c.String("u")

	if url == "" {
		fmt.Println("flag -u (url) is required")
		return
	}

	fmt.Println(url)
	n._host = url

	n._status = true
}

func (n *NETWORK) _getHandleBack() string {
	return ""
}

func (n *NETWORK) _initConfig(r string) {
	c := config{
		status: true,
		method: "GET",
	}

	c.headers = []string{"UserAgent"}
	request := r + n._getHandleBack()

	if n._method == 0 {
		c.url = n._host + "?" + n._parameter + "=" + request
	} else if n._method == 1 {
		c.form = request
	} else if n._method == 2 {
		c.headers = []string{request}
	} else if n._method == 3 {
		//cookie
		c.jar = []string{}
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
		fmt.Println("HERE")
		err := new(error)
		panic(err)
	}

	req, err := http.NewRequest(n._config.method, n._config.url, nil)

	if err != nil {
		fmt.Println("HERE2")
		panic(err)
	}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	n._handleSuccess(resp)
}
