package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func add(arg ...int) int {
	sum := 0
	for _, n := range arg {
		sum += n
	}

	return sum
}

func myIncr(n *int) {
	*n++
}

func add_two(n_slice []int) {
	for i := 0; i < len(n_slice); i++ {
		n_slice[i] += 2
	}
}

func main() {
	max_num := max(3, 10)
	fmt.Printf("max: %d\n", max_num)

	add_num := add(1, 2, 3)
	fmt.Printf("add: %v\n", add_num)

	n := 10
	myIncr(&n)
	fmt.Printf("n: %d\n", n)

	n_slice := []int{10, 20, 30}
	fmt.Printf("n_slice Before: %v\n", n_slice)
	add_two(n_slice)
	fmt.Printf("n_slice After: %v\n", n_slice)
}
