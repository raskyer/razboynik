package worker

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

type apidata struct {
	Config razboy.Config
	Action string
}

type apiresponse struct {
	Status   string
	Response string
}

func Api(port string) error {
	http.HandleFunc("/api/exec", _apiExec)

	return http.ListenAndServe(":"+port, nil)
}

func _apiExec(w http.ResponseWriter, req *http.Request) {
	var (
		k       *kernel.Kernel
		kc      *kernel.KernelCmd
		decoder *json.Decoder
		api     *apidata
		apires  apiresponse
		res     []byte
		err     error
	)

	api = new(apidata)

	decoder = json.NewDecoder(req.Body)
	err = decoder.Decode(api)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer req.Body.Close()

	k = kernel.Boot()
	kc = kernel.CreateCmd(api.Action)
	kc, err = k.Exec(kc, &api.Config)

	if err != nil {
		apires = apiresponse{
			Status:   "error",
			Response: err.Error(),
		}
	} else {
		apires = apiresponse{
			Status:   "success",
			Response: kc.GetResult(),
		}
	}

	res, err = json.Marshal(apires)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(res)
}
