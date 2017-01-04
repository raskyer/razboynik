package examplemodule

import "github.com/eatbytes/razboynik/services/kernel"
import "github.com/eatbytes/razboy"

type Fibocmd struct{}

func (f *Fibocmd) Exec(kl *kernel.KernelLine, config *razboy.Config) error {
	var (
		action  string
		request *razboy.REQUEST
		//response *razboy.RESPONSE
		err error
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
	_, err = razboy.Send(request)

	if err != nil {
		return err
	}

	//kl.WriteSuccess(response.GetResult())

	return nil
}

func (f *Fibocmd) GetName() string {
	return "-fibo"
}

func (f *Fibocmd) GetCompleter() (kernel.CompleterFunction, bool) {
	return nil, false
}
