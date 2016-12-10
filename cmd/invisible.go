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

	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var invisibleCmd = &cobra.Command{
	Use:   "invisible [url] [referer]",
	Short: "Send invisible request",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			result string
			err    error
		)

		if len(args) < 2 {
			return errors.New("not enough arguments")
		}

		if !silent {
			printer.PrintIntro()
			printer.PrintSection("Invisible", "Send invisible request")
		}

		result, err = worker.Invisible(args[0], args[1])

		if err != nil {
			return err
		}

		if !silent {
			printer.PrintSection("Result", result)
		} else {
			printer.Println(result)
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(invisibleCmd)
}
