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
	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [url]",
	Short: "Run reverse shell with configuration",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var c *razboy.Config

		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		if !silent {
			printer.PrintIntro()
			printer.PrintSection("Run", "Run reverse shell with configuration")
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

		return worker.Run(c)
	},
}

func init() {
	RootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&method, "method", "m", "GET", "Method to use. Ex: -m POST")
	runCmd.Flags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	runCmd.Flags().StringVarP(&key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	runCmd.Flags().StringVarP(&shellmethod, "shellmethod", "s", "system", "System function used in php script. Ex: -s shell_exec")
	runCmd.Flags().StringVarP(&encoding, "encoding", "e", "base64", "Encoding of the request. Ex: -e base64")
	runCmd.Flags().StringVar(&shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	runCmd.Flags().StringVar(&proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
