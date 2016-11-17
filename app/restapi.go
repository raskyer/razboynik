package app

import (
	"encoding/json"
	"net/http"

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services"
	"github.com/urfave/cli"
)

type apirequest struct {
	Cmd    string
	Scope  string
	Method int
}

type apidata struct {
	Config  core.Config
	Request apirequest
}

func (app *AppInterface) Api(c *cli.Context) {
	var port string

	port = c.String("p")

	services.PrintSection("API", "REST API launch on port : "+port)

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
		services.PrintError(err)
		return
	}
}

func (app *AppInterface) apiphp(w http.ResponseWriter, req *http.Request) {

}
