package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL = "https://www.google.com/"

func main() {
	resp, err := http.Get(baseURL)
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
