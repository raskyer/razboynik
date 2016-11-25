package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/normalizer"
)

func Invisible(url, referer string) (string, error) {
	var (
		request *razboy.REQUEST
		rzRes   *razboy.RazResponse
		err     error
	)

	request = razboy.CreateRequest(
		[4]string{url, "GET", "", ""},
		[2]string{"", ""},
		[2]bool{false, false},
	)

	request.Headers = []razboy.HEADER{
		razboy.HEADER{
			Key:   "Referer",
			Value: normalizer.Encode(referer),
		},
	}

	rzRes, err = razboy.Send(request)

	if err != nil {
		return "", err
	}

	return rzRes.GetResult(), nil
}
