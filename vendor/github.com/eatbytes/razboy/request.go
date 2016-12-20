package razboy

import "net/http"

type _shellscope struct {
	Name    string
	Content []string
}

type HEADER struct {
	Key   string
	Value string
}

type REQUEST struct {
	Action     string
	UploadPath string
	Upload     bool
	Headers    []HEADER
	body       []byte
	cmd        string
	setup      bool
	c          *Config
	http       *http.Request
}

func CreateRequest(action string, c *Config) *REQUEST {
	return &REQUEST{
		Action: action,
		c:      c,
	}
}

func (req REQUEST) IsProtected() bool {
	if req.c.Key != "" {
		return true
	}

	return false
}

func (req REQUEST) GetHTTP() *http.Request {
	return req.http
}

func (req REQUEST) GetConfig() *Config {
	return req.c
}

func (req REQUEST) GetBody() []byte {
	return req.body
}
