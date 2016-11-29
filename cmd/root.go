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
	"os"

	"github.com/eatbytes/razboynik/services/booter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	method      string
	parameter   string
	key         string
	shellmethod string
	encoding    string
	proxy       string
	shellscope  string
	raw         bool /* @DEPRECATED */
	debug       bool
	silent      bool
)

var RootCmd = &cobra.Command{
	Use:   "razboynik",
	Short: "Virtual shell on webserver thanks to file upload vulnerability",
	Long:  ``,
}

func Execute() {
	booter.Boot()

	if err := RootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&method, "method", "m", "GET", "Method to use. Ex: -m POST")
	RootCmd.PersistentFlags().StringVarP(&parameter, "parameter", "p", "razboynik", "Parameter to use. Ex: -p test")
	RootCmd.PersistentFlags().StringVarP(&key, "key", "k", "", "Key to unlock optional small protection. Ex: -k keytounlock")
	RootCmd.PersistentFlags().StringVarP(&shellmethod, "shellmethod", "s", "system", "System function used in php script. Ex: -s shell_exec")
	RootCmd.PersistentFlags().StringVarP(&encoding, "encoding", "e", "base64", "Encoding of the request. Ex: -e base64")
	RootCmd.PersistentFlags().StringVar(&shellscope, "scope", "", "Scope inside the shell. Ex: --scope /var")
	RootCmd.PersistentFlags().StringVar(&proxy, "proxy", "", "Proxy url where request will be sent before. Ex: --proxy http://localhost:8080")
	RootCmd.PersistentFlags().BoolVarP(&raw, "raw", "r", false, "*DEPRECATED* (use encoding instead) raw")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Print more information for debugging. Ex: --debug")
	RootCmd.PersistentFlags().BoolVar(&silent, "silent", false, "Don't print anything else than result or error. Ex: --silent")
}

func initConfig() {
	viper.SetConfigName(".razboynik")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
