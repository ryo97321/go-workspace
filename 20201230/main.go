package main

import "fmt"

func main() {
	p := struct {
		name string
		age  int
	}{
		name: "Gopher",
		age:  10,
	}

	p.age++

	fmt.Println(p.name)
	fmt.Println(p.age)
}
