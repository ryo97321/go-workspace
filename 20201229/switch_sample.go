package main

import "fmt"

func main() {
	var a int = 1

	switch {
	case a == 1:
		fmt.Println("a is 1 or 2")
	default:
		fmt.Println("default")
	}
}
