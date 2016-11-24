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
			err     error
			kc      *kernel.KernelCmd
			request *core.REQUEST
		)

		if len(args) < 2 {
			return errors.New("Not enough arguments.")
		}

		printer.PrintIntro()
		printer.PrintSection("Execute", "Execute a command on server")

		shl := worker.BuildShellConfig(shellmethod, "")
		php := worker.BuildPHPConfig(raw, false)
		srv := worker.BuildServerConfig(args[0], method, parameter, key, raw)

		request = worker.BuildRequest(shl, php, srv)

		kc, err = worker.Exec(args[1], request)

		if debug && kc.GetRzResp() != nil {
			debugger.HTTP(kc.GetRzResp())
		}

		if err != nil {
			return err
		}

		printer.PrintSection("Result", kc.GetResult())

		return nil
	},
}

func init() {
	RootCmd.AddCommand(execCmd)
}
