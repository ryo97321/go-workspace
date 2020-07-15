package main

import "fmt"

func main() {
	var a interface{}

	var i int = 5
	s := "Hello World"

	a = i
	fmt.Printf("a: %v\n", a)

	a = s
	fmt.Printf("a: %v\n", a)
}
