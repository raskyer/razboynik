package fuzzer

import "encoding/base64"

func Encode(str string) string {
	sEnc := base64.StdEncoding.EncodeToString([]byte(str))

	return sEnc
}

func Decode(str string) string {
	sDec, err := base64.StdEncoding.DecodeString(str)

	if err != nil {
		ferr := NormalizeErr(err)
		ferr.Error()
		return ""
	}

	return string(sDec)
}

func PHPEncode(str string) string {
	return "base64_encode(" + str + ")"
}

func PHPDecode(str string) string {
	return "base64_decode(" + str + ")"
}
