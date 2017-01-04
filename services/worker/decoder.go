package worker

import "github.com/eatbytes/razboy"

func Encode(str string) string {
	return razboy.Encode(str)
}

func Decode(str string) (string, error) {
	sDec, err := razboy.Decode(str)

	if err != nil {
		return str, err
	}

	return sDec, nil
}
