package main

import "fmt"

const (
	x = iota
	y = iota
	z = iota
)

func main() {
	fmt.Printf("x: %d, y: %d, z: %d\n", x, y, z)
}
