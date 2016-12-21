package worker

import (
	"bytes"
	"encoding/json"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/fatih/color"
)

type scanresult struct {
	S struct {
		Sy string
		Sh string
		Pr string
		Pa string
	}
	I struct {
		W string
		P string
	}
}

func test(n string) string {
	if n == "1" {
		return "-" + color.GreenString("[v]")
	}

	return "-" + color.RedString("[x]")
}

func Scan(config *razboy.Config) (string, error) {
	var (
		s       *scanresult
		kc      *kernel.KernelCmd
		decoder *json.Decoder
		m       []string
		info    [2]string
		result  string
		err     error
	)

	s = new(scanresult)
	m = []string{"GET", "POST", "HEADER", "COOKIE"}

	for i := 0; i < 4; i++ {
		result += "\nMethod: " + color.YellowString(m[i]) + "\n"

		config.Method = m[i]
		kc, err = Exec("-scan", config)

		if err != nil {
			result += "-" + color.RedString("[x]") + " Error Exec: " + err.Error() + "\n"
			continue
		}

		decoder = json.NewDecoder(bytes.NewReader([]byte(kc.GetResult())))
		err = decoder.Decode(&s)

		if err != nil {
			result += "-" + color.RedString("[x]") + " Error Decoder: " + err.Error() + "\n"
			continue
		}

		result += "-" + color.GreenString("[v]")
		result += " Mode: " + color.MagentaString("Raw PHP") + "\n"

		result += test(s.S.Sy)
		result += " Mode: " + color.MagentaString("SHELL system") + "\n"

		result += test(s.S.Sh)
		result += " Mode: " + color.MagentaString("SHELL shell_exec") + "\n"

		result += test(s.S.Pr)
		result += " Mode: " + color.MagentaString("SHELL proc_open") + "\n"

		result += test(s.S.Pa)
		result += " Mode: " + color.MagentaString("SHELL passthru") + "\n"

		if s.I.P != "" {
			info[1] = s.I.P
		}

		if s.I.W != "" {
			info[0] = s.I.W
		}
	}

	result += "\nInfo:\n"
	result += "User: " + info[0] + "\n"
	result += "pwd: " + info[1]

	return result, nil
}
