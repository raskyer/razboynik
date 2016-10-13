package php

import (
	"github.com/eatbytes/fuzzcore"
	"github.com/leaklessgfy/fuzzer/reader"
)

func Upload(path, dir string) {
	bytes, bondary, err := fuzzcore.PHP.Upload(path, dir)

	if err != nil {
		err.Error()
		return
	}

	req, err := fuzzcore.NET.PrepareUpload(bytes, bondary)

	if err != nil {
		err.Error()
		return
	}

	resp, err := fuzzcore.NET.Send(req)

	if err != nil {
		err.Error()
		return
	}

	result := fuzzcore.NET.GetBodyStr(resp)
	reader.ReadOne(result, "File upload successfully")
}
