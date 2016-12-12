package worker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type apidata struct {
	config razboy.Config
	action string
}

func Api(port string) {
	http.HandleFunc("/api/shell", _apishell)
	http.HandleFunc("/api/php", _apiphp)
	http.ListenAndServe(":"+port, nil)
}

func _apishell(w http.ResponseWriter, req *http.Request) {
	var (
		decoder *json.Decoder
		api     *apidata
		err     error
	)

	decoder = json.NewDecoder(req.Body)
	err = decoder.Decode(api)

	if err != nil {
		fmt.Println(err)
	}

	defer req.Body.Close()

	fmt.Println(api)

	k := kernel.Boot()
	kc := kernel.CreateCmd(api.action)

	kc, err = k.Exec(kc, &api.config)
}

func _apiphp(w http.ResponseWriter, req *http.Request) {

}
