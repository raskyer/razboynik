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

package target

import (
	"errors"

	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker"
	"github.com/spf13/cobra"
)

var EditCmd = &cobra.Command{
	Use:   "edit [target]",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		if cmd.Flag("silent").Value.String() == "false" {
			printer.PrintIntro()
			printer.PrintSection("Edit target", "Edit target '"+args[0]+"' in config file")
		}

		return worker.TargetEdit(args[0])
	},
}
