package main

import (
	"fmt"
	"strings"
)

func getTweetFullTextLines(fullText string) []string {
	fullTextLines := strings.Split(fullText, "\n")

	return fullTextLines
}

func main() {
	tweetFullText := "My\nname\nis\n***\n"

	tweetFullTextLines := getTweetFullTextLines(tweetFullText)
	for _, line := range tweetFullTextLines {
		fmt.Println(line)
	}
}
