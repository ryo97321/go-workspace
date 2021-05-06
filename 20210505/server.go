package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PostHelloRequest struct {
	Name  string `json:name`
	Times int    `json:times`
}

type PostHelloResponse struct {
	Hellos []string `json:"hellos"`
}

func postHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var reqBody PostHelloRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name, times := reqBody.Name, reqBody.Times

	hellos := make([]string, 0, times)
	for i := 0; i < times; i++ {
		hellos = append(hellos, "Hello "+name)
	}

	resBody := PostHelloResponse{hellos}
	encoder := json.NewDecoder(w)
	encoder.SetIndent("", "	")
	w.Header().Set("Content-Type", "application/json")
	if err := encoder.Encode(resBody); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

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
