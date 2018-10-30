# CMCCLI
CMCCLI stands for **C**oin **M**arket **C**ap **C**ommand **L**ine **I**nterface.

**Origin**: I started this project because I didn't have any knowledge of Go and wanted to learn it.

CMCCLI is built with Go and uses [cobra](https://github.com/spf13/cobra) for handling command line inputs, [gabs](https://github.com/Jeffail/gabs) for handling dynamic json data and [tablewriter](https://github.com/olekukonko/tablewriter) for displaying tables. CMCCLI is also licensed under [GPLv3](/LICENSE.md)

As the name suggests, CMCCLI is based on the [coinmarketcap](https://coinmarketcap.com) API.
The old public API will stop working on December 4th 2018. CMCCLI is built on the newly introduced [professional-api](https://pro.coinmarketcap.com) by coinmarketcap.

**DISCLAIMER**: CMCCLI is still very early development and many things are going to change, if you want to get involved, look at the issues tab and make a pull request.

## How to get started

The only way to use CMCCLI at the moment is to clone this repo and compile it yourself.

### Adding your API key
[make an account](https://pro.coinmarketcap.com) to get an API key.
Once you have your API key, run: 

`CMCCLI setApiKey [your key]`

### Setting conversion currency

Set whatever fiat currency you want your prices converted to. [Here](https://pro.coinmarketcap.com/api/v1#section/Standards-and-Conventions) is a full list of fiat currencies you can choose.

`CMCCLI setConversion [currency]`

Example:

`CMCCLI setConversion USD`

### Stats for a single currency

`CMCCLI coin [symbol of currency]`

Example:

`CMCCLI coin btc`

### List the top cryptocurrencies ranked by marketcap

`CMCCLI list [n]` n = how many to list (default 100)

**every 200 returned data points are 1 call credit (rounded up) so be careful to not list too many**

### Watch command

`CMCCLI watch add [symbol]`

example:

`CMCCLI watch add BTC`

then run:

`CMCCLI watch` to list all currencies you saved to your portfolio.

### Portfolio command

to add something to your portfolio:

`CMCCLI port buy [symbol] [price]`

`CMCCLI port buy btc 1`

Output:

`total portfolio worth: 6460.89410671 USD` 

to remove something again:

`CMCCLI port rm [symbol]`