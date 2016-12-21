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
	"os"

	"github.com/eatbytes/razboynik/services/booter"
	"github.com/spf13/cobra"
)

var (
	method      string
	parameter   string
	key         string
	shellmethod string
	encoding    string
	proxy       string
	shellscope  string
	debug       bool
	silent      bool
	noextra     bool
)

var RootCmd = &cobra.Command{
	Use:   "razboynik",
	Short: "Reverse shell on webserver thanks to file upload vulnerability. version: 2.0.0",
	Long:  ``,
}

func Execute() {
	booter.Boot()

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "Don't print anything else than result or error. Ex: --silent")
}
