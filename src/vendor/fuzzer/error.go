package fuzzer

import "fmt"

type FuzzerError struct {
	code int
	msg  string
	bag  error
	conf *config
}

func (e FuzzerError) Error() string {
	fmt.Printf("%v %v \n", e.code, e.msg)
	fmt.Printf("bag: %v \n", e.bag)
	fmt.Printf("conf: %v \n", e.conf)
	return fmt.Sprintf("%v: %v - { bag: %v } { conf: %v }", e.code, e.msg, e.bag, e.conf)
}

func SetupErr() FuzzerError {
	return FuzzerError{
		msg: "Error: You havn't setup the required information, please refer to srv config.",
	}
}

func RequestErr(err error, c int) FuzzerError {
	return FuzzerError{
		code: c,
		msg:  "Error: Impossible to send request.",
		bag:  err,
	}
}

func BuildRequestErr(err error, c *config) FuzzerError {
	return FuzzerError{
		msg:  "Error: Impossible to create request with config",
		bag:  err,
		conf: c,
	}
}

func NoMethodFoundErr() FuzzerError {
	return FuzzerError{
		msg: "Error: No method was find for the req to prepare",
	}
}

func FileErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Encounter a problem with file",
		bag: err,
	}
}

func PartErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Can't create part",
		bag: err,
	}
}

func NormalizeErr(err error) FuzzerError {
	return FuzzerError{
		msg: "Error: Impossible to normalize the string",
		bag: err,
	}
}
