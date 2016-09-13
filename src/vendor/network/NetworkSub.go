package network

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/urfave/cli"
)

func SpecifiedItem(c *cli.Context) {
	fmt.Println("Please specified the item: response or request")
}

func showRequestUrl(r *http.Request) {
	fmt.Println(r.URL)
}

func showRequestMethod(r *http.Request) {
	fmt.Println(r.Method)
}

func showRequestBody(r *http.Request) {
	r.ParseForm()
	fmt.Println(r)
	fmt.Println(r.Body)
}

func showRequestHeaders(r *http.Request) {
	fmt.Println(r.Header)
}

func showResponseStatus(r *http.Response) {
	fmt.Println(r.Status)
}

func showResponseBody(r *http.Response) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("body: %v", string(body))
}

func showResponseHeaders(r *http.Response) {
	fmt.Println(r.Header)
}

func showResponseRequest(r *http.Response) {
	fmt.Println(r.Request)
}
