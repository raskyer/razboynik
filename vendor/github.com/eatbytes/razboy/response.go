package razboy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/eatbytes/razboy/normalizer"
)

type RESPONSE struct {
	http    *http.Response
	request *REQUEST
}

func (res *RESPONSE) GetHTTP() *http.Response {
	return res.http
}

func (res *RESPONSE) GetRequest() *REQUEST {
	return res.request
}

func (res *RESPONSE) GetBody() []byte {
	var (
		buffer []byte
		err    error
	)

	defer res.http.Body.Close()

	buffer, err = ioutil.ReadAll(res.http.Body)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	res.http.Body = ioutil.NopCloser(bytes.NewReader(buffer))

	return buffer
}

func (res *RESPONSE) GetBodyStr() string {
	return string(res.GetBody())
}

func (res *RESPONSE) GetHeaderStr() string {
	return res.http.Header.Get(res.request.c.Parameter)
}

func (res *RESPONSE) GetCookieStr() string {
	var (
		str     string
		cookies []*http.Cookie
		cookie  *http.Cookie
	)

	cookies = res.http.Cookies()

	for _, cookie = range cookies {
		if cookie.Name == res.request.c.Parameter {
			str, _ = url.QueryUnescape(cookie.Value)
			return str
		}
	}

	return ""
}

func (res *RESPONSE) GetResultStrByMethod(m string) string {
	if m == "HEADER" {
		return res.GetHeaderStr()
	}

	if m == "COOKIE" {
		return res.GetCookieStr()
	}

	return res.GetBodyStr()
}

func (res *RESPONSE) GetResultStr() string {
	return res.GetResultStrByMethod(res.request.c.Method)
}

func (res *RESPONSE) GetResult() string {
	var str string

	str = res.GetResultStr()
	str, _ = normalizer.Decode(str)

	return str
}
