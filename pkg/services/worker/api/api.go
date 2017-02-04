package api

import (
	"encoding/json"
	"net/http"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/pkg/services/kernel"
	"github.com/eatbytes/razboynik/pkg/services/worker/configuration"
	"github.com/eatbytes/razboynik/pkg/services/worker/printer"
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
		printer.PrintError(err)
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

func getResponse() []byte {
	var (
		res    apiresponse
		buffer []byte
		err    error
	)

	if err != nil {
		res = apiresponse{
			Status:   "error",
			Response: "resp.Err.Error()",
		}
	} else {
		res = apiresponse{
			Status:   "success",
			Response: "resp.Body.(string)",
		}
	}

	buffer, err = json.Marshal(res)

	if err != nil {
		printer.PrintError(err)
		return nil
	}

	return buffer
}

func apiExec(w http.ResponseWriter, req *http.Request) {
	var (
		k      *kernel.Kernel
		config *razboy.Config
		d      *data
		res    []byte
	)

	d = extractData(req)
	config = getConfig(d)

	k = kernel.Boot()
	err := k.Exec(d.Action, config)

	printer.PrintError(err)
	res = getResponse()
	w.Write(res)
}
