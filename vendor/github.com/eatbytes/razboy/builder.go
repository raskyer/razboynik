package razboy

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func _createSimpleRequest(req *REQUEST) error {
	var err error

	req.setup = true
	req.cmd = Encode(req.Action)

	switch req.c.Method {
	case "GET":
		err = _buildGET(req)
		break
	case "POST":
		err = _buildPOST(req)
		break
	case "HEADER":
		err = _buildHEADER(req)
		break
	case "COOKIE":
		err = _buildCOOKIE(req)
		break
	default:
		req.setup = false
	}

	if err != nil {
		return err
	}

	_addHeaders(req)
	err = _addProxy(req)

	return err
}

func _createUploadRequest(req *REQUEST) error {
	var (
		writer *multipart.Writer
		file   *os.File
		body   *bytes.Buffer
		part   io.Writer
		err    error
	)

	req.setup = true
	req.cmd = Encode(req.Action)

	file, err = os.Open(req.UploadPath)

	if err != nil {
		return err
	}

	defer file.Close()

	body = &bytes.Buffer{}
	writer = multipart.NewWriter(body)
	part, err = writer.CreateFormFile("file", filepath.Base(req.UploadPath))

	if err != nil {
		return err
	}

	_, err = io.Copy(part, file)

	if err != nil {
		return err
	}

	writer.WriteField(req.c.Parameter, req.cmd)

	if req.IsProtected() {
		writer.WriteField(KEY, req.c.Key)
	}

	err = writer.Close()

	if err != nil {
		return err
	}

	req.http, err = http.NewRequest("POST", req.c.Url, body)

	if err != nil {
		return err
	}

	req.body = body.Bytes()
	req.http.Header.Add("Content-Type", writer.FormDataContentType())

	_addHeaders(req)

	return _addProxy(req)
}

func _addHeaders(req *REQUEST) {
	if len(req.Headers) > 0 {
		for _, header := range req.Headers {
			req.http.Header.Add(header.Key, header.Value)
		}
	}
}

func _addProxy(req *REQUEST) error {
	var (
		proxy *url.URL
		err   error
	)

	if req.c.Proxy != "" {
		proxy, err = url.Parse(req.c.Proxy)

		if err != nil {
			return err
		}

		http.DefaultTransport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	}

	return nil
}

func _buildGET(req *REQUEST) error {
	var (
		url string
		err error
	)

	url = req.c.Url + "?" + req.c.Parameter + "=" + req.cmd

	if req.IsProtected() {
		url += "&" + KEY + "=" + req.c.Key
	}

	req.http, err = http.NewRequest("GET", url, nil)

	return err
}

func _buildPOST(req *REQUEST) error {
	var (
		form url.Values
		data *bytes.Buffer
		err  error
	)

	form = url.Values{}
	form.Set(req.c.Parameter, req.cmd)

	if req.IsProtected() {
		form.Add(KEY, req.c.Key)
	}

	data = bytes.NewBufferString(form.Encode())
	req.http, err = http.NewRequest("POST", req.c.Url, data)

	if err != nil {
		return err
	}

	req.body = data.Bytes()
	req.http.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	return nil
}

func _buildHEADER(req *REQUEST) error {
	var err error

	req.http, err = http.NewRequest("GET", req.c.Url, nil)

	if err != nil {
		return err
	}

	req.http.Header.Add(req.c.Parameter, req.cmd)

	if req.IsProtected() {
		req.http.Header.Add(KEY, req.c.Key)
	}

	return nil
}

func _buildCOOKIE(req *REQUEST) error {
	var (
		cookie, kcookie *http.Cookie
		err             error
	)

	req.http, err = http.NewRequest("GET", req.c.Url, nil)

	if err != nil {
		return err
	}

	cookie = &http.Cookie{Name: req.c.Parameter, Value: req.cmd}
	req.http.AddCookie(cookie)

	if req.IsProtected() {
		kcookie = &http.Cookie{Name: KEY, Value: req.c.Key}
		req.http.AddCookie(kcookie)
	}

	return nil
}
