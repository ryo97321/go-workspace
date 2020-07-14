package main

import "fmt"

func sum(nums ...int) (result int) {
	for _, n := range nums {
		result += n
	}
	return
}

func main() {
	fmt.Printf("sum: %d\n", sum(1, 2, 3, 4))
}
