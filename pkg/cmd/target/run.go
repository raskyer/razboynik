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

package target

import (
	"errors"

	"github.com/eatbytes/razboynik/pkg/services/kernel"
	"github.com/eatbytes/razboynik/pkg/services/worker/configuration"
	"github.com/eatbytes/razboynik/pkg/services/worker/printer"
	"github.com/eatbytes/razboynik/pkg/services/worker/target"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run [target]",
	Short: "Run reverse shell on target in config file",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		var (
			k      *kernel.Kernel
			config *configuration.Configuration
			target *target.Target
			err    error
		)

		if len(args) < 1 {
			return errors.New("not enough arguments")
		}

		printer.PrintIntro()
		printer.PrintSection("Run target", "Run target '"+args[0]+"' in config file")

		config, err = configuration.GetConfiguration()

		if err != nil {
			return err
		}

		target, _, err = configuration.FindTarget(config, args[0])

		if err != nil {
			return err
		}

		k = kernel.Boot()

		return k.Run(target.Config)
	},
}
