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
	"strconv"
	"strings"
	"time"
)

var (
	cur Currency
)

type Currency struct {
	Data struct {
		BTC struct {
			ID                int         `json:"id"`
			Name              string      `json:"name"`
			Symbol            string      `json:"symbol"`
			Slug              string      `json:"slug"`
			CirculatingSupply float64     `json:"circulating_supply"`
			TotalSupply       float64     `json:"total_supply"`
			MaxSupply         interface{} `json:"max_supply"`
			DateAdded         time.Time   `json:"date_added"`
			NumMarketPairs    int         `json:"num_market_pairs"`
			CmcRank           int         `json:"cmc_rank"`
			LastUpdated       time.Time   `json:"last_updated"`
			Quote             struct {
				USD struct {
					Price            float64   `json:"price"`
					Volume24H        float64   `json:"volume_24h"`
					PercentChange1H  float64   `json:"percent_change_1h"`
					PercentChange24H float64   `json:"percent_change_24h"`
					PercentChange7D  float64   `json:"percent_change_7d"`
					MarketCap        float64   `json:"market_cap"`
					LastUpdated      time.Time `json:"last_updated"`
				} `json:"USD"`
			} `json:"quote"`
		} //`json:"BTC"`
	} `json:"data"`
}

// coinCmd represents the coin command
var coinCmd = &cobra.Command{
	Use:   "coin",
	Short: "coin [symbol of coin] like BTC or ETH (upper or lowercase)",
	Long: `All cryptocurrencies have a symbol/short name like BTC or ETH for example. When giving this value to the coin command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Price: " + Price(args))
	},
}

func Price(args []string) (string) {
	resp := gv.GetFromApi("/cryptocurrency/quotes/latest?symbol=" + strings.ToUpper(strings.Join(args, "")))
	jsonParsed, err := gabs.ParseJSON(resp)
	if err != nil {
		log.Fatal(err)
	}

	var price float64
	var ok bool
	symbol := strings.Join(args, "")
	symbol = strings.ToUpper(symbol)
	fmt.Println(symbol)

	price, ok = jsonParsed.Search("data", symbol, "quote", "USD", "price").Data().(float64)
	if ok {
		return strconv.FormatFloat(price, 'f', -1, 64)
	}
	return ""
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
