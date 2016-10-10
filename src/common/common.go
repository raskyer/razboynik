package common

import (
	"fmt"
	"net/http"

	"github.com/eatbytes/fuzzcore"
	"github.com/eatbytes/sysgo"
)

func Send(str string) (*http.Response, error) {
	req, err := fuzzcore.NET.Prepare(str)

	if err != nil {
		return nil, err
	}

	resp, err := fuzzcore.NET.Send(req)

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

	result := fuzzcore.NET.GetResultStr(resp)

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
	sDec := fuzzcore.Decode(str)
	fmt.Println(sDec)
}

func Read(str string) {
	fmt.Println(str)
}

func Syscall(str string) {
	result, err := sysgo.Call(str)
	
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
