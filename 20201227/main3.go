package main

import "fmt"

func main() {
	var price int
	fmt.Print("値段>")
	fmt.Scan(&price)
	fmt.Printf("%d円\n", price)
}
