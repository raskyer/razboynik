package worker

import "github.com/eatbytes/razboy"

func Invisible(url, referer string) (string, error) {
	var (
		c        *razboy.Config
		request  *razboy.REQUEST
		response *razboy.RESPONSE
		err      error
	)

	c = &razboy.Config{
		Url:       url,
		Method:    "GET",
		Parameter: "",
		Key:       "",
	}

	request = razboy.CreateRequest("", c)
	request.Headers = []razboy.HEADER{
		razboy.HEADER{
			Key:   "Referer",
			Value: razboy.Encode(referer),
		},
	}

	response, err = razboy.Send(request)

	if err != nil {
		return "", err
	}

	return response.GetResult(), nil
}
