package worker

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboy/normalizer"
)

func Invisible(url, referer string) (string, error) {
	var (
		request *core.REQUEST
		rzRes   *razboy.RazResponse
		err     error
	)

	request = &core.REQUEST{
		Type: "PHP",
		SRVc: &core.SERVERCONFIG{
			Url:    url,
			Method: "GET",
			Headers: []core.HEADER{
				core.HEADER{Key: "Referer", Value: normalizer.Encode(referer)},
			},
		},
	}

	rzRes, err = razboy.Send(request)

	if err != nil {
		return "", err
	}

	return rzRes.GetResult(), nil
}
