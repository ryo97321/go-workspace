package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "https://api.bitflyer.jp/v1/getticker?product_code="

func main() {
	resp, err := http.Get(baseURL + "btc_jpy")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArray))
}
