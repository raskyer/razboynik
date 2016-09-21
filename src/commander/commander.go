package commander

import "fuzzer"

func Process(str string) (string, bool) {
	req, err := fuzzer.NET.Prepare(str)

	if err {
		return "", true
	}

	resp, err := fuzzer.NET.Send(req)

	if err {
		return "", true
	}

	result := fuzzer.NET.GetBodyStr(resp)

	return result, false
}
