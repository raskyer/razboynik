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
	"github.com/eatbytes/razboynik/services/debugger"
	"github.com/eatbytes/razboynik/services/kernel"
	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec [url] [cmd]",
	Short: "A brief description of your command",
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
			Raw:         raw,
			Proxy:       proxy,
			Shellmethod: shellmethod,
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

	execCmd.LocalFlags().StringVarP(&method, "method", "m", "GET", "Method to use. Ex: -m POST")
	execCmd.LocalFlags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	execCmd.LocalFlags().StringVarP(&key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	execCmd.LocalFlags().StringVarP(&shellmethod, "shellmethod", "s", "system", "System function used in php script. Ex: -s shell_exec")
	execCmd.LocalFlags().StringVarP(&encoding, "encoding", "e", "base64", "Encoding of the request. Ex: -e base64")
	execCmd.LocalFlags().StringVar(&shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	execCmd.LocalFlags().BoolVarP(&raw, "raw", "r", false, "*DEPRECATED* (use encoding instead) raw")
	execCmd.LocalFlags().BoolVar(&debug, "debug", false, "Print more information for debugging. Ex: --debug")
	execCmd.LocalFlags().StringVar(&proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
