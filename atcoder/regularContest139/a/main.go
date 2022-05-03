package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	if n > 20 {
		fmt.Println("upper 20")
	} else {
		fmt.Println("down 20")
	}
}
