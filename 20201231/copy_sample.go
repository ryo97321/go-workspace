package main

import "fmt"

func main() {
	p := struct {
		age  int
		name string
	}{
		age:  10,
		name: "Gopher",
	}

	p2 := p

	p2.age = 100
	p2.name = "Gopher2"

	fmt.Println(p)
	fmt.Println(p2)
}
