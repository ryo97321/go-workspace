package main

import "fmt"

func main() {
	var msg string

	// for i := 1; i <= 100; i++ {
	// 	if i%2 == 0 {
	// 		msg = "偶数"
	// 	} else if i%2 == 1 {
	// 		msg = "奇数"
	// 	}

	// 	fmt.Printf("%d-%s\n", i, msg)
	// }

	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0:
			msg = "偶数"
		case i%2 == 1:
			msg = "奇数"
		default:
			msg = ""
		}

		fmt.Printf("%d-%s\n", i, msg)
	}
}
