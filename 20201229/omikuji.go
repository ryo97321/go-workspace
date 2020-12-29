package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	value := rand.Intn(7)

	var result string

	if value == 1 {
		result = "凶"
	} else if value == 2 || value == 3 {
		result = "吉"
	} else if value == 4 || value == 5 {
		result = "中吉"
	} else if value == 6 {
		result = "大吉"
	}

	fmt.Println(result)
}
