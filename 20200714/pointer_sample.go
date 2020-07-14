package main

import "fmt"

func callByValue(i int) {
	i = 20
}

func callByRef(i *int) {
	*i = 20
}

func main() {
	var i int = 10

	callByValue(i)
	fmt.Printf("callByValue: %d\n", i)

	callByRef(&i)
	fmt.Printf("callByRef: %d\n", i)
}
