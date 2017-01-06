package razboy

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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

func (res *RESPONSE) GetRawBody() []byte {
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

func (res *RESPONSE) GetRawBodyStr() string {
	return string(res.GetRawBody())
}

func (res *RESPONSE) GetRawHeaderStr() string {
	return res.http.Header.Get(res.request.c.Parameter)
}

func (res *RESPONSE) GetRawCookieStr() string {
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

func (res *RESPONSE) GetRawResultStrByMethod(m int) string {
	if m == M_HEADER {
		return res.GetRawHeaderStr()
	}

	if m == M_COOKIE {
		return res.GetRawCookieStr()
	}

	return res.GetRawBodyStr()
}

func (res *RESPONSE) GetRawResultStr() string {
	return res.GetRawResultStrByMethod(res.request.c.Method)
}

func (res *RESPONSE) GetResult() string {
	var str string

	str = res.GetRawResultStr()
	str, _ = Decode(str)

	return str
}
