// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"

	"github.com/eatbytes/razboy"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan [url]",
	Short: "Scan target by url and display possible method to use",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			c   *razboy.Config
			kc  *kernel.KernelCmd
			err error
		)

		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		if !silent {
			printer.PrintIntro()
			printer.PrintSection("Scan", "Scan target by url and display possible method to use")
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

		kc, err = worker.Scan(c)

		if err != nil {
			return err
		}

		if !silent {
			printer.PrintTitle("Result")
		}

		printer.Println(kc.GetResult())

		return nil
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	scanCmd.Flags().StringVarP(&key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	scanCmd.Flags().StringVar(&shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	scanCmd.Flags().StringVar(&proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
