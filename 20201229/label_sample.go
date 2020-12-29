package main

import "fmt"

func main() {
	var i int

LOOP:
	for {
		switch {
		case i%2 == 1:
			break LOOP
		}
		i++
	}

	fmt.Printf("i = %d\n", i)
}
