package common

import (
	"bytes"
	"fmt"
	"fuzzer"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func Process(str string) (string, bool) {
	req, err := fuzzer.NET.Prepare(str)

	if err {
		return "", true
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return "", true
	}

	result := fuzzer.NET.GetResultStr(resp)

	return result, false
}

func Upload(path, dir string) {
	bytes, bondary, err := fuzzer.PHP.Upload(path, dir)

	if err {
		return
	}

	req, err := fuzzer.NET.PrepareUpload(bytes, bondary)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)

	ReadOne(result, "File upload successfully")
}

func Download(path, location string) {
	php := fuzzer.PHP.Download(path)
	req, err := fuzzer.NET.Prepare(php)

	if err {
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return
	}

	ReadDownload(resp, location)
}

func ReadDownload(resp *http.Response, location string) {
	out, err := os.Create(location)
	defer out.Close()

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Downloaded successfully those byte: ")
	fmt.Println(resp.Body)
}

func ReadOne(r, msg string) {
	if r == "1" {
		fmt.Println(msg)
		return
	}

	fmt.Println("An error occured")
}

func ReadEncode(str string) {
	sDec := fuzzer.Decode(str)
	fmt.Println(sDec)
}

func Read(str string) {
	fmt.Println(str)
}

func Syscall(str string) {
	cmd := exec.Command("bash", "-c", str)

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	err := cmd.Run()

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err))
	}

	fmt.Printf("%s\n", string(cmdOutput.Bytes()))
}
