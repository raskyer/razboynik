package app

import (
	"net/http"

	"github.com/urfave/cli"
)

func (app *AppInterface) Api(c *cli.Context) {
	var port string

	port = c.String("p")

	http.HandleFunc("/api/shell", app.apishell)
	http.HandleFunc("/api/php", app.apiphp)
	http.ListenAndServe(":"+port, nil)
}

func (app *AppInterface) apishell(w http.ResponseWriter, req *http.Request) {

}

func (app *AppInterface) apiphp(w http.ResponseWriter, req *http.Request) {

}
