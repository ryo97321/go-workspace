package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Employee struct {
	Human
	speciality string
	phone      string
}

func main() {
	Bob := Employee{Human{"Bob", 34, "777-444-XXXX"}, "Designer", "333-222"}
	fmt.Printf("Bob phone: %s\n", Bob.phone)
	fmt.Printf("Bob Human phone: %s\n", Bob.Human.phone)
}
