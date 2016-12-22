// package phpmodule

// import (
// 	"errors"
// 	"os"
// 	"os/exec"

// 	"io/ioutil"

// 	"strings"

// 	"io"

// 	"github.com/eatbytes/razboy"
// 	"github.com/eatbytes/razboynik/services/kernel"
// )

// func PHPScriptStatic(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
// 	var (
// 		request *razboy.REQUEST
// 		buffer  []byte
// 		bufstr  string
// 		err     error
// 	)

// 	if kc.GetArrLgt() < 2 {
// 		return kc, errors.New("Please write the path of the local file to upload")
// 	}

// 	buffer, err = ioutil.ReadFile(kc.GetArrItem(1))

// 	if err != nil {
// 		return kc, err
// 	}

// 	bufstr = string(buffer)
// 	bufstr = strings.TrimPrefix(bufstr, "<php")

// 	request = razboy.CreateRequest(bufstr, c)
// 	_, err = kc.Send(request)

// 	return kc, err
// }

// func PHPScriptInteractive(kc *kernel.KernelCmd, c *razboy.Config) (*kernel.KernelCmd, error) {
// 	var (
// 		cmd     *exec.Cmd
// 		request *razboy.REQUEST
// 		buffer  []byte
// 		bufstr  string
// 		local   string
// 		err     error
// 	)

// 	local = "/tmp/tmp-razboynik.php"

// 	if kc.GetArrLgt() > 1 {
// 		f, err := os.Create(local)

// 		if err != nil {
// 			return kc, err
// 		}

// 		f2, err := os.OpenFile(kc.GetArrItem(1), os.O_RDONLY, 0666)

// 		if err != nil {
// 			return kc, err
// 		}

// 		_, err = io.Copy(f, f2)
// 	}

// 	cmd = exec.Command("vim", local)
// 	cmd.Stdin = os.Stdin
// 	cmd.Stdout = os.Stdout
// 	err = cmd.Run()

// 	if err != nil {
// 		return kc, err
// 	}

// 	buffer, err = ioutil.ReadFile(local)

// 	if err != nil {
// 		return kc, err
// 	}

// 	bufstr = string(buffer)
// 	bufstr = strings.TrimPrefix(bufstr, "<php")

// 	request = razboy.CreateRequest(bufstr, c)
// 	_, err = kc.Send(request)

// 	return kc, err
// }
