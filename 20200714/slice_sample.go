package main

import "fmt"

func main() {
	var s []string
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c")

	for _, elem := range s {
		fmt.Println(elem)
	}
}
