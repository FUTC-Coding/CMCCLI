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
	"github.com/fatih/color"

	//"CMCCLI/gv"
	"fmt"
	"github.com/Jeffail/gabs"
	"log"
	"os"
	"strconv"
	"strings"

	//"strings"

	"github.com/spf13/cobra"
)

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "with port you can list your saved portfolio",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		showPort()
	},
}

var cmdBuy = &cobra.Command{
	Use:   "buy",
	Short: "port buy [symbol] [amount]",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		amount, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			fmt.Println("please add an amount")
			log.Fatal(err)
		}
		if args[0] == "" {
			fmt.Println("please add the symbol of the currency and the amount you own, to the buy command")
			os.Exit(-3)
		} else {
			buy(args[0], amount)
		}
	},
}

var cmdRm = &cobra.Command{
	Use:   "rm",
	Short: "port rm [symbol]",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("please add the currency you want to remove, example: btc")
			os.Exit(-3)
		}
		remove(strings.Join(args, ""))
	},
}

func showPort(){
	//parse portfolio.json file with gabs
	jsonFile, err := gabs.ParseJSONFile("portfolio.json")
	if err != nil {
		log.Fatal(err)
	}

	var sym []string
	slice := sym
	//append all the symbol names to a string array for the api request
	children, _ := jsonFile.S("symbol").Children()
	for _, child := range children {
		slice = append(slice, strings.ToUpper(child.Data().(string))) //make all symbols uppercase because the api requires the symbols to be uppercase
	}

	var am []float64
	sliced := am
	//append all the amounts to a float64 array to calculate the total worth of the portfolio
	children_a, _ := jsonFile.S("amount").Children()
	for _, child := range children_a {
		sliced = append(sliced, child.Data().(float64))
	}
	//get data from api here with string array of symbols
	prices := getPricesFromApi(slice)
	//calculate total worth of portfolio with array of amount and array of price
	total := calcProfits(sliced, prices)
	color.Cyan("total portfolio worth: " + strconv.FormatFloat(total, 'f', -1, 64) + " USD")
}


func getPricesFromApi(symbols []string) ([]float64) {
	s := strings.Join(symbols, ",")
	//get and parse the json data from the api
	resp := gv.GetFromApi("/cryptocurrency/quotes/latest?symbol=" + s)
	jsonParsed, err := gabs.ParseJSON(resp)
	if err != nil {
		log.Fatal(err)
	}

	var prices []float64
	slice := prices
	i := 0
	//append the price of each currency that was fetched to a float64 array
	for range symbols {
		price,_ := jsonParsed.S("data", symbols[i], "quote", "USD","price").Data().(float64)
		slice = append(slice, price)
		i++
	}
	return slice
}

func calcProfits(amounts []float64, prices []float64) (float64){
	var total float64
	a := len(amounts)
	for i := 0 ; i < a; i++{
		total = total + amounts[i] * prices[i]
	}
	return total
}

func buy(symbol string, amount float64){
	if _, err := os.Stat("portfolio.json"); !os.IsNotExist(err) { //if file already exists, overwrite it with new old and new data merged
		jsonData, err := gabs.ParseJSONFile("portfolio.json")
		if err != nil {
			log.Fatal(err)
		}
		jsonData.ArrayAppend(symbol, "symbol")
		jsonData.ArrayAppend(amount, "amount")
		writePortFile(jsonData)
	} else { //if file doesn't exist yet, create it and write to it
		jsonObj := gabs.New()
		jsonObj.Array("symbol")
		jsonObj.Array("amount")
		jsonObj.ArrayAppend(symbol ,"symbol")
		jsonObj.ArrayAppend(amount ,"amount")
		writePortFile(jsonObj)
	}
}

func remove(symbol string) {
	//parse portfolio.json file with gabs
	jsonFile, err := gabs.ParseJSONFile("portfolio.json")
	if err != nil {
		log.Fatal(err)
	}
	var index int

	children,_ := jsonFile.S("symbol").Children()
	for _, child := range children {
		if child.Data().(string) == symbol {
			jsonFile.ArrayRemove(index, "symbol")
			jsonFile.ArrayRemove(index, "amount")
			fmt.Println("removed " + symbol + " from your portfolio")
			writePortFile(jsonFile)
			os.Exit(-3)
		} else {
			index++
		}
	}
	fmt.Println(symbol + " is not in your portfolio")
}

func writePortFile(jsonData *gabs.Container){
	if _, err := os.Stat("portfolio.json"); os.IsNotExist(err) {
		fmt.Println("file does not exist and cant be overwritten")
		os.Exit(-3)
	} else { //write file
		f, err := os.Create("portfolio.json")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		fmt.Fprintln(f, jsonData.StringIndent("", "  "))
	}
}

func init() {
	rootCmd.AddCommand(portCmd)
	portCmd.AddCommand(cmdBuy)
	portCmd.AddCommand(cmdRm)
}
