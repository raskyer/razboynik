package fuzzer

var FORMATER = FORMATERInterface{}

type FORMATERInterface struct{}

func (f *FORMATERInterface) Response() string {
	var response string

	m := NET.GetMethod()
	p := NET.GetParameter()

	if m == 0 || m == 1 {
		response = "echo(" + PHPEncode("$r") + ");exit();"
	} else if m == 2 {
		response = "header('" + p + ":' . " + PHPEncode("$r") + ");exit();"
	} else if m == 3 {
		response = "setcookie('" + p + "', " + PHPEncode("$r") + ");exit();"
	}

	return response
}
