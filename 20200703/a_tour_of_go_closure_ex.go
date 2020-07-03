package main

import "fmt"

func fibonacci() func() int {
	a1 := 0
	a2 := 1
	count := 0

	return func() int {
		count++
		if count == 1 {
			return 0
		} else if count == 2 {
			return 1
		} else {
			temp := a2
			a2 = a1 + a2
			a1 = temp
			return a2
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
