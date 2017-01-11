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
	"github.com/eatbytes/razboynik/cmd/target"
	"github.com/eatbytes/razboynik/services/worker/configuration"
	"github.com/eatbytes/razboynik/services/worker/printer"
	"github.com/spf13/cobra"
)

var targetCmd = &cobra.Command{
	Use:   "target",
	Short: "Target handler",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			config *configuration.Configuration
			err    error
		)

		printer.PrintIntro()

		config, err = configuration.GetConfiguration()

		if err != nil {
			return err
		}

		printer.PrintTitle("TARGETS")
		for _, item := range config.Targets {
			printer.Println("- " + item.Name)
		}
		printer.Print("\n")

		return nil
	},
}

func init() {
	targetCmd.AddCommand(target.AddCmd)
	targetCmd.AddCommand(target.EditCmd)
	targetCmd.AddCommand(target.RemoveCmd)
	targetCmd.AddCommand(target.RunCmd)
	targetCmd.AddCommand(target.DetailCmd)
	RootCmd.AddCommand(targetCmd)
}
