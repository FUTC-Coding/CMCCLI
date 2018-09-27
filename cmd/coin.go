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
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"log"
	"math"
	"strconv"
	"strings"
)

type Currency struct {
	Name string
	Rank float64
	CirculatingSupply float64
	MaxSupply float64
	Price float64
	Volume24H float64
	PercentChange1H float64
	PercentChange24H float64
	PercentChange7D float64
	MarketCap float64
}

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "coin [symbol of coin] like BTC or ETH (upper or lowercase)",
	Long: `All cryptocurrencies have a symbol/short name like BTC or ETH for example. When giving this value to the coin command`,
	Run: func(cmd *cobra.Command, args []string) {
		AllData(args)
	},
}

func AllData(args []string) {
	c := Currency{}
	symbol := strings.Join(args, "")
	symbol = strings.ToUpper(symbol)

	resp := gv.GetFromApi("/cryptocurrency/quotes/latest?symbol=" + symbol)
	jsonParsed, err := gabs.ParseJSON(resp)
	if err != nil {
		log.Fatal(err)
	}

	c.Name = jsonParsed.Search("data", symbol, "name").Data().(string)
	color.Magenta(c.Name)

	c.Rank = jsonParsed.Search("data", symbol, "cmc_rank").Data().(float64)
	fmt.Print("Rank: ")
	color.Yellow(strconv.FormatFloat(c.Rank, 'f', -1, 64))

	c.Price = jsonParsed.Search("data", symbol, "quote", "USD", "price").Data().(float64)
	fmt.Print("Price: ")
	color.Cyan(strconv.FormatFloat(math.Round(c.Price*1000000)/1000000, 'f', -1, 64) + " USD") //print in color blue; convert the format to a string; round to 6 decimals after the comma

	c.Volume24H = jsonParsed.Search("data", symbol, "quote", "USD", "volume_24h").Data().(float64)
	fmt.Print("Volume 24h: ")
	color.Cyan(strconv.FormatFloat(math.Round(c.Volume24H*1000000)/1000000, 'f', -1, 64) + " USD")

	c.PercentChange1H = jsonParsed.Search("data", symbol, "quote", "USD", "percent_change_1h").Data().(float64)
	x := strconv.FormatFloat(math.Round(c.PercentChange1H*1000000)/1000000, 'f', -1, 64) + "%"
	fmt.Print("1h change: ")
	if !Abs(c.PercentChange1H){
		color.Red(x)
	} else {
		color.Green(x)
	}

	c.PercentChange24H = jsonParsed.Search("data", symbol, "quote", "USD", "percent_change_24h").Data().(float64)
	x = strconv.FormatFloat(math.Round(c.PercentChange24H*1000000)/1000000, 'f', -1, 64) + "%"
	fmt.Print("24h change: ")
	if !Abs(c.PercentChange24H){
		color.Red(x)
	} else {
		color.Green(x)
	}

	c.PercentChange7D = jsonParsed.Search("data", symbol, "quote", "USD", "percent_change_7d").Data().(float64)
	x = strconv.FormatFloat(math.Round(c.PercentChange7D*1000000)/1000000, 'f', -1, 64) + "%"
	fmt.Print("7d change: ")
	if !Abs(c.PercentChange24H){
		color.Red(x)
	} else {
		color.Green(x)
	}

	c.MarketCap = jsonParsed.Search("data", symbol, "quote", "USD", "market_cap").Data().(float64)
	fmt.Print("Market Cap: ")
	color.Cyan(strconv.FormatFloat(math.Round(c.MarketCap*1000000)/1000000, 'f', -1, 64) + " USD")

	c.CirculatingSupply = jsonParsed.Search("data", symbol, "circulating_supply").Data().(float64)
	fmt.Print("Circulating Supply: ")
	color.Cyan(strconv.FormatFloat(c.CirculatingSupply, 'f', -1, 64))

	c.MaxSupply = jsonParsed.Search("data", symbol, "max_supply").Data().(float64)
	fmt.Print("Max. Supply: ")
	color.Cyan(strconv.FormatFloat(c.MaxSupply, 'f', -1, 64))
}

//this function checks if a float is positive or negative; it returns true if it is positive
func Abs(x float64) bool {
	if x > 0 {
		return true
	}
	return false
}

func init() {
	rootCmd.AddCommand(coinCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// coinCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// coinCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
