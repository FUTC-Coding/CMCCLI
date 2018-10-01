# CMCCLI
CMCCLI stands for **C**oin **M**arket **C**ap **C**ommand **L**ine **I**nterface.

**Origin**: I started this project because I didn't have any knowledge of Go and wanted to learn it.

CMCCLI is built with Go and uses [cobra](https://github.com/spf13/cobra) for handling command line inputs and [gabs](https://github.com/Jeffail/gabs) for handling dynamic json data. CMCCLI is also licensed under [GPLv3](/LICENSE.md)

As the name suggests, CMCCLI is based on the [coinmarketcap](https://coinmarketcap.com) API.
The old public API will stop working on December 4th 2018. CMCCLI is built on the newly introduced [professional-api](https://pro.coinmarketcap.com) by coinmarketcap.

**DISCLAIMER**: CMCCLI is still very early development and many things are going to change, if you want to get involved, look at the issues tab and make a pull request.

## How to get started

The only way to use CMCCLI at the moment is to clone this repo and compile it yourself.

### Adding your API key
[make an account](https://pro.coinmarketcap.com) to get an API key.
Once you have your API key, run: 

`CMCCLI setApiKey [your key]`

### Stats for a single currency

`CMCCLI coin [symbol of currency]`

example:

`CMCCLI coin btc`

### List the top cryptocurrencies ranked by marketcap

`CMCCLI list [n]` n = how many to list (default 100)

**every 100 returned data points are 1 call credit (rounded up) so be careful to not list too many**

### Portfolio command

`CMCCLI port add [symbol]`

example:

`CMCCLI port add BTC`

then run:

`CMCCLI port` to list all currencies you saved to your portfolio.
