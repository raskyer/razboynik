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

	"github.com/eatbytes/razboy/core"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var method string
var parameter string

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec [url] [cmd]",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			url    string
			raw    string
			config *core.SERVERCONFIG
		)

		if len(args) < 2 {
			return errors.New("Not enough arguments.")
		}

		url = args[0]
		raw = args[1]

		config = &core.SERVERCONFIG{
			Url:       url,
			Method:    method,
			Parameter: parameter,
		}

		worker.Exec(raw, config)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(execCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	execCmd.Flags().StringVarP(&method, "method", "m", "GET", "Method to use")
	execCmd.Flags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use")

}
