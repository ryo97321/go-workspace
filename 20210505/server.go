package main

import (
	"log"
	"net/http"
)

func getHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Hello, World!"))
}

func getGoodMorning(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Good Morning!"))
}

func main() {
	http.HandleFunc("/", getHello)
	http.HandleFunc("/morning/", getGoodMorning)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
