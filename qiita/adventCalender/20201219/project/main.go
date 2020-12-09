package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func calcAnimalAgeInHumanYears(animalLife, animalAge float64) float64 {
	humanLife := 84.1
	animalAgeInHumanYears := animalAge * (humanLife / animalLife)
	return animalAgeInHumanYears
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, nil)
}

func calcAnimalAgeInHumanYearsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	animalAgeStr := r.Form["animalAge"][0]
	animalLifeStr := r.Form["animalLife"][0]

	animalAge, _ := strconv.ParseFloat(animalAgeStr, 64)
	animalLife, _ := strconv.ParseFloat(animalLifeStr, 64)
	animalAgeInHumanYears := calcAnimalAgeInHumanYears(animalLife, animalAge)

	tpl := template.Must(template.ParseFiles("calcAnimalAgeInHumanYears.html"))
	tpl.Execute(w, animalAgeInHumanYears)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/calcAnimalAgeInHumanYears", calcAnimalAgeInHumanYearsHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
