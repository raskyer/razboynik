package razboy

const MAX_SIZE = 512

const (
	M_GET    = 0
	M_POST   = 1
	M_HEADER = 2
	M_COOKIE = 3

	SM_SHELL_EXEC = 0
	SM_SYSTEM     = 1
	SM_PASSTHRU   = 2

	E_BASE64 = 0

	PARAMETER = "razboynik"
	KEY       = "RAZBOYNIK_KEY"
)

type Config struct {
	Url         string
	Method      int
	Parameter   string
	Key         string
	Proxy       string
	Encoding    int
	Shellmethod int
	Shellscope  string
}

func NewConfig() *Config {
	return &Config{
		Method:      M_GET,
		Parameter:   PARAMETER,
		Shellmethod: SM_SHELL_EXEC,
		Encoding:    E_BASE64,
	}
}

func MethodToInt(m string) int {
	switch m {
	case "GET":
		return M_GET
	case "POST":
		return M_POST
	case "HEADER":
		return M_HEADER
	case "COOKIE":
		return M_COOKIE
	default:
		return M_GET
	}
}

func ShellmethodToInt(sm string) int {
	switch sm {
	case "shell_exec":
		return SM_SHELL_EXEC
	case "system":
		return SM_SYSTEM
	case "passthru":
		return SM_PASSTHRU
	default:
		return SM_SHELL_EXEC
	}
}

func MethodToStr(m int) string {
	switch m {
	case M_GET:
		return "GET"
	case M_POST:
		return "POST"
	case M_HEADER:
		return "HEADER"
	case M_COOKIE:
		return "COOKIE"
	default:
		return "GET"
	}
}

func ShellmethodToStr(sm int) string {
	switch sm {
	case SM_SHELL_EXEC:
		return "shell_exec"
	case SM_SYSTEM:
		return "system"
	case SM_PASSTHRU:
		return "passthru"
	default:
		return "shell_exec"
	}
}
