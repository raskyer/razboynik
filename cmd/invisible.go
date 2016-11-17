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
			url, referer, result string
			err                  error
		)

		printer.PrintIntro()
		printer.PrintSection("Invisible", "Send invisible request")

		if len(args) < 1 {
			return errors.New("Not enough arguments. First argument should be the url of the target")
		}

		if len(args) < 2 {
			return errors.New("Not enough arguments. Second argument should be the referer")
		}

		url = args[0]
		referer = args[1]

		result, err = worker.Invisible(url, referer)

		if err != nil {
			return err
		}

		printer.PrintSection("Result", result)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(invisibleCmd)
}
