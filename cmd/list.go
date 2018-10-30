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
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"strings"
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
	conversion := gv.ReadConversion()
	if len(args) == 0 { //if no args, list 100
		resp := gv.GetFromApi("/cryptocurrency/listings/latest?start=1&limit=100&convert=" + conversion)
		jsonParsed, err := gabs.ParseJSON(resp)
		if err != nil {
			log.Fatal(err)
		}

		list(jsonParsed)

	} else if _, err := strconv.Atoi(a); err != nil { //check if args are numeric
		fmt.Print("please enter a valid number")
	} else { //list with given number
		resp := gv.GetFromApi("/cryptocurrency/listings/latest?start=1&limit=" + a + "&convert=" + conversion)
		jsonParsed, err := gabs.ParseJSON(resp)
		if err != nil {
			log.Fatal(err)
		}

		list(jsonParsed)
	}
}

func list(jsonParsed *gabs.Container) {
	conversion := gv.ReadConversion()
	table := tablewriter.NewWriter(os.Stdout)
	price := "Price " + conversion
	mc := "Marketcap  " + conversion
	volume := "Volume 24h " + conversion
	table.SetHeader([]string{"Rank", "Name", price, mc, "Circulating supply", "Total supply", volume, "Change 1h (%)", "Change 24h (%)", "Change 7d (%)"})
	table.SetBorder(false)

	var data [][]string
	sliced := data

	names, _ := jsonParsed.S("data").Children()
	for _, child := range names { //iterate through all currencies and output rank and name

		//get all data from the json and save in variables and convert to string for output
		rank := strconv.FormatFloat(child.Search("cmc_rank").Data().(float64), 'f', -1, 64)
		name := child.Search("name").Data().(string)
		price := strconv.FormatFloat(child.Search("quote", conversion, "price").Data().(float64), 'f', -1, 64)
		marketcap := strconv.FormatFloat(child.Search("quote", conversion, "market_cap").Data().(float64),'f',-1,64)
		circulating := strconv.FormatFloat(child.Search("circulating_supply").Data().(float64),'f',-1,64)
		total := strconv.FormatFloat(child.Search("total_supply").Data().(float64),'f',-1,64)
		volume24 := strconv.FormatFloat(child.Search("quote", conversion, "volume_24h").Data().(float64),'f',-1,64)

		//get the change rates but keep as float to check if they are positive or negative
		change1 := child.Search("quote", conversion, "percent_change_1h").Data().(float64)
		change24 := child.Search("quote", conversion, "percent_change_24h").Data().(float64)
		change7 := child.Search("quote", conversion, "percent_change_7d").Data().(float64)

		//convert change rates to string for output
		change1s := strconv.FormatFloat(change1,'f',-1,64)
		change24s := strconv.FormatFloat(change24,'f',-1,64)
		change7s := strconv.FormatFloat(change7,'f',-1,64)
		var a,b,c string
		if !Abs(change1){ //if change 1h is positive make it green. If it's negative, make it red
			a = "\033[31m" + change1s + "\033[0m"
		} else {
			a = "\033[32m" + change1s + "\033[0m"
		}

		if !Abs(change24){ //if change 24h is positive make it green. If it's negative, make it red
			b = "\033[31m" + change24s + "\033[0m"
		} else {
			b = "\033[32m" + change24s + "\033[0m"
		}

		if !Abs(change7){ //if change 7d is positive make it green. If it's negative, make it red
			c = "\033[31m" + change7s + "\033[0m"
		} else {
			c = "\033[32m" + change7s + "\033[0m"
		}

		//append to the output
		sliced = append(sliced, []string{rank,name,price,marketcap,circulating,total,volume24,a,b,c})
	}
	table.SetHeaderColor(tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgCyanColor}) //set Header color to cyan for all strings in the header

	table.SetColumnColor(tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{},
		tablewriter.Colors{}) //set color of the rank column to yellow

	table.AppendBulk(sliced)
	table.Render() //actually output the table
}

func init() {
	rootCmd.AddCommand(listCmd)
}
