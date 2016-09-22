package common

import (
	"bytes"
	"fmt"
	"fuzzer"
	"net/http"
	"os"
	"os/exec"
)

func Send(str string) (*http.Response, error) {
	req, err := fuzzer.NET.Prepare(str)

	if err != nil {
		return nil, err
	}

	resp, err := fuzzer.NET.Send(req)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func Process(str string) (string, error) {
	resp, err := Send(str)

	if err != nil {
		return "", err
	}

	result := fuzzer.NET.GetResultStr(resp)

	return result, nil
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
