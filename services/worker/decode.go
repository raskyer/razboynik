package worker

import "github.com/eatbytes/razboy/normalizer"

func Encode(str string) string {
	return normalizer.Encode(str)
}

func Decode(str string) (string, error) {
	sDec, err := normalizer.Decode(str)

	if err != nil {
		return str, err
	}

	return sDec, nil
}
