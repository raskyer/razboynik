package worker

import (
	"io/ioutil"
	"net/http"
)

func GetBody(r *http.Response) []byte {
	defer r.Body.Close()
	buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	return buffer
}
