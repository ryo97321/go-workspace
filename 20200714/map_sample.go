package main

import "fmt"

func main() {
	month := map[int]string{
		1: "January",
		2: "February",
	}

	for key, value := range month {
		fmt.Printf("%d: %s\n", key, value)
	}
}
