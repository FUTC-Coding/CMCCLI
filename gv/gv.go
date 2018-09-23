package gv

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Api struct {
	key string
}

var apikey string

func (a *Api) SetApiKey(key string) {
	//for debugging
	fmt.Println("apikey is parameter: " + a.key)

	a.key = key

	//for debugging
	fmt.Println("apikey has been set: " + a.key)

}

func (a Api) ApiKey() (string) {
	fmt.Println("getting apikey: " + a.key)
	return a.key
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readKey() (string){
	dat, err := ioutil.ReadFile(".apikey")
	check(err)
	fmt.Print("read this from file: " + string(dat))
	return string(dat)
}

func GetFromApi(directory string) {

	const baseURL = "https://pro-api.coinmarketcap.com/v1"
	url := fmt.Sprintf(baseURL + directory)

	cmcClient := http.Client{
		Timeout: time.Second * 2,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	//for debugging
	fmt.Println("requesting with apikey: " + readKey())

	request.Header.Set("X-CMC_PRO_API_KEY", readKey())

	response, err := cmcClient.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}