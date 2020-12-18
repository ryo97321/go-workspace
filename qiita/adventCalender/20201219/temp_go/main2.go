package main

import "fmt"

func main() {
	const fastSmalls = true
	const nSmalls = 100
	var base int = 10

	var i int = -10

	fmt.Println(fastSmalls && 0 <= i)
	fmt.Println(i < nSmalls && base == 10)
	fmt.Println(fastSmalls && 0 <= i && i < nSmalls && base == 10)
}
