package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "https://api.bitflyer.jp/v1/getticker?product_code="

type Response struct {
	ProductCode string  `json:"product_code"`
	LTP         float64 `json:"ltp"`
}

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

	mapRes := new(Response)
	if err := json.Unmarshal(byteArray, &mapRes); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("プロダクト名: %+v\n", mapRes.ProductCode)
	fmt.Printf("価格: %+v\n", mapRes.LTP)
}
