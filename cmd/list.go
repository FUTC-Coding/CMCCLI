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
	"CMCCLI/gv"
	"fmt"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list [n]; n = number of currencies to list, default 100.",
	Long: `To list a specific amount of currencies, use list [number]. The default number of currencies that are being listed, is 100`,
	Run: func(cmd *cobra.Command, args []string) {
		listAll(args)
	},
}

func listAll(args []string) {
	a := strings.Join(args, "")



	if len(args) == 0 { //if no args, list 100
		resp := gv.GetFromApi("/cryptocurrency/listings/latest?start=1&limit=100&convert=USD")
		jsonParsed, err := gabs.ParseJSON(resp)
		if err != nil {
			log.Fatal(err)
		}

		list(jsonParsed)

	} else if _, err := strconv.Atoi(a); err != nil { //check if args are numeric
		fmt.Print("please enter a valid number")
	} else { //list with given number
		resp := gv.GetFromApi("/cryptocurrency/listings/latest?start=1&limit=" + a + "&convert=USD")
		jsonParsed, err := gabs.ParseJSON(resp)
		if err != nil {
			log.Fatal(err)
		}

		list(jsonParsed)
	}
}

func list(jsonParsed *gabs.Container) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)

	fmt.Fprintln(w, "Rank\tName\tPrice (USD)\tMarket Cap\tCirculating Supply\tTotal Supply\tVolume 24h\tChange 1h\tChange 24h\tChange 7d")

	names, _ := jsonParsed.S("data").Children() //iterate through all currencies and output rank and name
	for _, child := range names {

		rank := strconv.FormatFloat(child.Search("cmc_rank").Data().(float64), 'f', -1, 64)
		name := child.Search("name").Data().(string)
		price := strconv.FormatFloat(child.Search("quote", "USD", "price").Data().(float64), 'f', -1, 64)
		marketcap := strconv.FormatFloat(child.Search("quote", "USD", "market_cap").Data().(float64),'f',-1,64)
		circulating := strconv.FormatFloat(child.Search("circulating_supply").Data().(float64),'f',-1,64)
		total := strconv.FormatFloat(child.Search("total_supply").Data().(float64),'f',-1,64)
		volume24 := strconv.FormatFloat(child.Search("quote", "USD", "volume_24h").Data().(float64),'f',-1,64)
		change1 := strconv.FormatFloat(child.Search("quote", "USD", "percent_change_1h").Data().(float64),'f',-1,64)
		change24 := strconv.FormatFloat(child.Search("quote", "USD", "percent_change_24h").Data().(float64),'f',-1,64)
		change7 := strconv.FormatFloat(child.Search("quote", "USD", "percent_change_7d").Data().(float64),'f',-1,64)

		fmt.Fprint(w, rank)
		fmt.Fprint(w, "\t")
		fmt.Fprintf(w, name  + "\t")
		fmt.Fprint(w, price + "\t")
		fmt.Fprint(w, marketcap + "\t")
		fmt.Fprint(w, circulating + "\t")
		fmt.Fprint(w, total + "\t")
		fmt.Fprint(w, volume24 + "\t")
		fmt.Fprint(w, change1 + "\t")
		fmt.Fprint(w, change24 + "\t")
		fmt.Fprintln(w, change7 + "\t")

	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
