package main

import "fmt"

func main() {
	s := "火曜日はTuesdayです。"
	runes := []rune(s) // []string -> []rune

	for _, r := range runes {
		c := string(r) // rune -> string
		fmt.Println(c)
	}
}
