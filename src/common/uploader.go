package common

import "fuzzer"

func Upload(path, dir string) {
	bytes, bondary, err := fuzzer.PHP.Upload(path, dir)

	if err != nil {
		err.Error()
		return
	}

	req, err := fuzzer.NET.PrepareUpload(bytes, bondary)

	if err != nil {
		err.Error()
		return
	}

	resp, err := fuzzer.NET.Send(req)

	if err != nil {
		err.Error()
		return
	}

	result := fuzzer.NET.GetBodyStr(resp)

	ReadOne(result, "File upload successfully")
}
