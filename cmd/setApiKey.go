// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"os"
	"strings"
)

// setApiKeyCmd represents the setApiKey command
var setApiKeyCmd = &cobra.Command{
	Use:   "setApiKey \"[apikey]\"",
	Short: "set the api key you can get from pro.coinmarketcap.com",
	Long: `This program makes API calls to the coinmarketcap api in the background,
to do this, it is necessary to have an API key. This API key is obtainable with a free account
at pro.coinmarketcap.com`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		writeKeyToFile(strings.Join(args,""))
	},
}

func init() {
	rootCmd.AddCommand(setApiKeyCmd)


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setApiKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setApiKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func writeKeyToFile(key string) {
	f, err := os.Create(".apikey")
	check(err)
	defer f.Close()

	fmt.Fprint(f, key)
}