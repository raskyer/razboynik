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
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/eatbytes/razboynik/services/worker/gflag"
	"github.com/eatbytes/razboynik/services/worker/printer"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [url]",
	Short: "Scan target by url and display possible method to use",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			c   *razboy.Config
			res string
			err error
		)

		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		printer.PrintIntro()
		printer.PrintSection("Scan", "Scan target by url and display possible method to use")

		c = gflag.BuildConfig(args[0])
		res, err = worker.Scan(c)

		if err != nil {
			return err
		}

		printer.PrintTitle("Result")
		printer.Println(res)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&gflag.Parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	scanCmd.Flags().StringVarP(&gflag.Key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	scanCmd.Flags().StringVar(&gflag.Shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	scanCmd.Flags().StringVar(&gflag.Proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
