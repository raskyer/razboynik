package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/eatbytes/razboy/core"
	"github.com/urfave/cli"
)

type apirequest struct {
	cmd    string
	scope  string
	method int
}

type apidata struct {
	config  core.Config
	request apirequest
}

func (app *AppInterface) Api(c *cli.Context) {
	var port string

	port = c.String("p")

	http.HandleFunc("/api/shell", app.apishell)
	http.HandleFunc("/api/php", app.apiphp)
	http.ListenAndServe(":"+port, nil)
}

func (app *AppInterface) apishell(w http.ResponseWriter, req *http.Request) {
	var (
		decoder *json.Decoder
		data    apidata
		err     error
	)

	decoder = json.NewDecoder(req.Body)
	err = decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	log.Println(data.request.cmd)
}

func (app *AppInterface) apiphp(w http.ResponseWriter, req *http.Request) {

}
