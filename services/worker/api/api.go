package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/worker/configuration"
)

type data struct {
	Target string
	Config razboy.Config
	Action string
}

type apiresponse struct {
	Status   string
	Response string
}

func Api(port string) error {
	http.HandleFunc("/api/exec", apiExec)

	return http.ListenAndServe(":"+port, nil)
}

func extractData(req *http.Request) *data {
	var (
		d       *data
		decoder *json.Decoder
		err     error
	)

	defer req.Body.Close()

	d = new(data)
	decoder = json.NewDecoder(req.Body)
	err = decoder.Decode(d)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return d
}

func getConfig(d *data) *razboy.Config {
	if d.Target != "" {
		c, err := configuration.GetConfiguration()

		if err != nil {
			return &d.Config
		}

		target, _, err := configuration.FindTarget(c, d.Target)

		if err != nil {
			return &d.Config
		}

		return target.Config
	}

	return &d.Config
}

func getResponse(resp kernel.Response) []byte {
	var (
		res    apiresponse
		buffer []byte
		err    error
	)

	if resp.Err != nil {
		res = apiresponse{
			Status:   "error",
			Response: resp.Err.Error(),
		}
	} else {
		res = apiresponse{
			Status:   "success",
			Response: resp.Body.(string),
		}
	}

	buffer, err = json.Marshal(res)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer
}

func apiExec(w http.ResponseWriter, req *http.Request) {
	var (
		k      *kernel.Kernel
		config *razboy.Config
		resp   kernel.Response
		d      *data
		res    []byte
	)

	d = extractData(req)
	config = getConfig(d)

	k = kernel.Boot()
	kernel.Silent()

	resp = k.Exec(d.Action, config)

	res = getResponse(resp)
	w.Write(res)
}
