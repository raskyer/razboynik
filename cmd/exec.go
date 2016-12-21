// MIT License

// Copyright (c) 2016

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cmd

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/debugger"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec [url] [cmd]",
	Short: "Execute a single command on url",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			err error
			kc  *kernel.KernelCmd
			c   *razboy.Config
		)

		if len(args) < 2 {
			return errors.New("not enough arguments")
		}

		if !silent {
			printer.PrintIntro()
			printer.PrintSection("Execute", "Execute a command on server")
		}

		c = &razboy.Config{
			Url:         args[0],
			Method:      method,
			Parameter:   parameter,
			Key:         key,
			Proxy:       proxy,
			Encoding:    encoding,
			Shellmethod: shellmethod,
			Shellscope:  shellscope,
		}

		kc, err = worker.Exec(args[1], c)

		if debug && kc.GetResponse() != nil {
			debugger.HTTP(kc.GetResponse())
		}

		if err != nil {
			return err
		}

		if !silent {
			printer.PrintSection("Result", kc.GetResult())
		} else {
			printer.Println(kc.GetResult())
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	execCmd.Flags().StringVarP(&method, "method", "m", "GET", "Method to use. Ex: -m POST")
	execCmd.Flags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	execCmd.Flags().StringVarP(&key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	execCmd.Flags().StringVarP(&shellmethod, "shellmethod", "s", "system", "System function used in php script. Ex: -s shell_exec")
	execCmd.Flags().StringVarP(&encoding, "encoding", "e", "base64", "Encoding of the request. Ex: -e base64")
	execCmd.Flags().StringVar(&shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	execCmd.Flags().BoolVar(&debug, "debug", false, "Print more information for debugging. Ex: --debug")
	execCmd.Flags().StringVar(&proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
