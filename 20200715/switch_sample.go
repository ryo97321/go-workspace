package main

import "fmt"

func main() {
	n := 1

	switch n {
	case 10:
		fmt.Println("n is 10")
	default:
		fmt.Println("n isn't 10")
	}

	count := 3

	switch count {
	case 3:
		fmt.Println("3")
		fallthrough
	case 2:
		fmt.Println("2")
		fallthrough
	case 1:
		fmt.Println("1")
		fallthrough
	default:
		fmt.Println("Boom!!")
	}
}
