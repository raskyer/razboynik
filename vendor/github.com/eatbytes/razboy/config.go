package razboy

type Config struct {
	Url         string
	Method      string
	Parameter   string
	Key         string
	Proxy       string
	Encoding    string
	Shellmethod string
	Shellscope  string
}

func NewConfig() *Config {
	c := new(Config)
	c.Method = "GET"
	c.Parameter = PARAM
	c.Shellmethod = "system"
	c.Encoding = "base64"

	return c
}
