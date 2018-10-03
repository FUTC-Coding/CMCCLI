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
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

var(
	symbol string
)

// listCmd represents the port command
var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "list your watched currencies",
	Long: `With just "list" you can list your saved currencies. With "list add [symbol of coin e.g. BTC]". With "list rm [BTC] you can remove it again from your watchlist"`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		listWatched()
	},
}

var cmdAdd = &cobra.Command{
	Use:   "add [BTC]",
	Short: "add a currency to the watchlist",
	Long: `Add a currency to your personal watchlist by supplying the symbol of the currency, e.g. BTC.`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := strings.Join(args, "")
		add(symbol)
		fmt.Println("added " + symbol + " to your watchlist")
	},
}

func listWatched(){
	if _, err := os.Stat(".watchlist"); !os.IsNotExist(err) { //if file already exists
		// open portfolio file
		f, err := os.Open(".watchlist")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var sa []string
			s := scanner.Text()
			sa = strings.Split(s, "")
			GetCoinData(sa)
			fmt.Println()
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else { //if file doesn't exist yet
		fmt.Println("You first have to add something to your watchlist. Use \"watch add [BTC] to add a currency to your watchlist\"")
	}
}

func add(symbol string) {
	if _, err := os.Stat(".watchlist"); !os.IsNotExist(err) { //if file already exists

		// open watchlist file
		f, err := os.OpenFile(".watchlist", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		w := bufio.NewWriter(f)

		fmt.Fprintln(w, symbol)

		w.Flush()
	} else { //if file doesn't exist yet, create it and write to it
		f, err := os.Create(".watchlist")
		check(err)
		defer f.Close()

		fmt.Fprintln(f, symbol)
	}

}

func init() {
	rootCmd.AddCommand(watchCmd)
	cmdAdd.Flags().StringVarP(&symbol, "add", "a", "", "add a currency to the personal watchlist")
	watchCmd.AddCommand(cmdAdd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
