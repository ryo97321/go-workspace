package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{name: "Bob", age: 20})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)
	fmt.Println(s.age)

	sp := &s

	fmt.Println(sp.age, s.age)
	sp.age = 500
	fmt.Println(sp.age, s.age)
}
