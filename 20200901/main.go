package main

import "fmt"

func main() {
	const (
		a = 1 << iota
		b
		c
		d
		e
		f
		g
	)

	fmt.Printf("a: %d, b: %d, c: %d, d: %d, e: %d, f: %d, g: %d\n", a, b, c, d, e, f, g)
}
