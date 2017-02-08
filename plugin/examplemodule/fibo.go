package examplemodule

import (
	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
)

var Fiboitem = kernel.Item{
	Name: "-fibo",
	Exec: func(l *kernel.Line, config *razboy.Config) kernel.Response {
		var (
			action   string
			request  *razboy.REQUEST
			response *razboy.RESPONSE
			err      error
		)

		action = `function fibo($n) {
			if($n < 2) {
				return 1;
			}

			return fibo($n - 1) + fibo($n - 2);
		}

		$r = array();
		for($i = 0; $i < 20; $i++) {
			$r[] = fibo($i);
		}

		$r = implode("\n", $r);`

		action += razboy.AddAnswer(config.Method, config.Parameter)
		request = razboy.CreateRequest(action, config)
		response, err = razboy.Send(request)

		if err != nil {
			return kernel.Response{Err: err}
		}

		kernel.WriteSuccess(l.GetStdout(), response.GetResult())

		return kernel.Response{Body: response.GetResult()}
	},
}
