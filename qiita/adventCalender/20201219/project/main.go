package main

import (
	"html/template"
	"net/http"
)

func calcAnimalAgeInHumanYears(humanLife, animalLife, animalAge float64) float64 {
	animalAgeInHumanYears := animalAge * (humanLife / animalLife)
	return animalAgeInHumanYears
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("index.html"))
	tpl.Execute(w, nil)
}

type params struct {
	AnimalAge  string
	AnimalLife string
}

func calcAnimalAgeInHumanYearsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	animalAge := r.Form["animalAge"][0]
	animalLife := r.Form["animalLife"][0]

	params := params{AnimalAge: animalAge, AnimalLife: animalLife}

	tpl := template.Must(template.ParseFiles("calcAnimalAgeInHumanYears.html"))
	tpl.Execute(w, params)
}

func main() {
	// humanLife := 80.0
	// animalLife := 16.0
	// animalAge := 1.0

	// animalAgeInHumanYears := calcAnimalAgeInHumanYears(humanLife, animalLife, animalAge)
	// fmt.Println(animalAgeInHumanYears)

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/calcAnimalAgeInHumanYears", calcAnimalAgeInHumanYearsHandler)
	port := "8080"
	http.ListenAndServe(":"+port, nil)
}
