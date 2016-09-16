package fuzzer

import (
	"io/ioutil"
	"net/http"
)

func GetBody(r *http.Response) []byte {
	if NET._respBody != nil {
		return NET._respBody
	}

	defer r.Body.Close()
	buffer, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	NET._respBody = buffer

	return buffer
}
