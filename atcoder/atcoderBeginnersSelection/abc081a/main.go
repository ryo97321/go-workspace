package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s", &s)

	var nBall int = 0
	for _, c := range s {
		if c == '1' {
			nBall++
		}
	}

	fmt.Println(nBall)
}
