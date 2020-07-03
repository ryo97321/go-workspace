package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(s)

	for _, word := range words {
		if _, isExist := wordMap[word]; isExist {
			wordMap[word]++
		} else {
			wordMap[word] = 1
		}
	}

	return wordMap
}

func main() {
	wc.Test(WordCount)
}
