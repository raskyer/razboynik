package examplemodule

import "github.com/eatbytes/razboynik/services/kernel"
import "github.com/eatbytes/razboy"

type Fibocmd struct{}

func (f *Fibocmd) Exec(kl *kernel.KernelLine, config *razboy.Config) kernel.KernelResponse {
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
		return kernel.KernelResponse{Err: err}
	}

	kernel.WriteSuccess(kl.GetStdout(), response.GetResult())

	return kernel.KernelResponse{Body: response.GetResult()}
}

func (f *Fibocmd) GetName() string {
	return "-fibo"
}

func (f *Fibocmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
