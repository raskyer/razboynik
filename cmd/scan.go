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
	"fmt"

	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan [url]",
	Short: "Scan a target by url",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("scan called")
	},
}

func init() {
	RootCmd.AddCommand(scanCmd)

	scanCmd.Flags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	scanCmd.Flags().StringVarP(&key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	scanCmd.Flags().StringVar(&shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	scanCmd.Flags().StringVar(&proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
}
