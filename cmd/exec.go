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
	"github.com/eatbytes/razboynik/services/gflags"
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
			kr kernel.Response
			c  *razboy.Config
		)

		if len(args) < 2 {
			return errors.New("not enough arguments")
		}

		printer.PrintIntro()
		printer.PrintSection("Execute", "Execute a command on server")

		c = gflags.BuildConfig(args[0])
		kr = worker.Exec(args[1], c)

		// if debug && kc.GetResponse() != nil {
		// 	debugger.HTTP(kc.GetResponse())
		// }

		if kr.Err != nil {
			return kr.Err
		}

		printer.PrintTitle("Result")
		printer.Println(kr.Body.(string))

		return nil
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	execCmd.Flags().StringVarP(&gflags.Method, "method", "m", "GET", "Method to use. Ex: -m POST")
	execCmd.Flags().StringVarP(&gflags.Parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	execCmd.Flags().StringVarP(&gflags.Key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	execCmd.Flags().StringVarP(&gflags.Shellmethod, "shellmethod", "s", "system", "System function used in php script. Ex: -s shell_exec")
	execCmd.Flags().IntVarP(&gflags.Encoding, "encoding", "e", razboy.E_BASE64, "Encoding of the request. Ex: -e 0")
	execCmd.Flags().StringVar(&gflags.Shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	execCmd.Flags().BoolVar(&gflags.Debug, "debug", false, "Print more information for debugging. Ex: --debug")
	execCmd.Flags().StringVar(&gflags.Proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
