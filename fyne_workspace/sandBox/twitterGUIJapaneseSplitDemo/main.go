package main

import (
	"fmt"
	"strings"
)

// fullTextを分割する（改行区切り かつ 1行の最大文字数は30文字）
func splitFullText(fullText string) []string {
	fullTextLines := strings.Split(fullText, "\n")

	var splitedFullText []string
	for _, fullTextLine := range fullTextLines {
		var splitedLines []string
		runes := []rune(fullTextLine)
		if len(runes) < 30 { // 30文字未満なら、そのままsplitedFullTextに加える
			splitedLines = append(splitedLines, string(runes))
		} else {
			s := ""
			for i, r := range runes {
				s += string(r)

				if (i+1)%30 == 0 {
					splitedLines = append(splitedLines, s)
					s = ""
				}
			}
			if s != "" { // s に文字列が残っていたらsplitedLinesに加える
				splitedLines = append(splitedLines, s)
			}
		}

		// 30文字ずつに分割した文字列スライスをsplitedFullTextに加える
		for _, splitedLine := range splitedLines {
			splitedFullText = append(splitedFullText, splitedLine)
		}
	}

	return splitedFullText
}

func main() {
	fullText := "月曜日はMondayです。火曜日はTuesdayです。水曜日はWednesdayです。木曜日はThuradayです。\n金曜日はFridayです。"

	fullTextLines := splitFullText(fullText)
	for _, fullTextLine := range fullTextLines {
		fmt.Printf("%s : %d\n", fullTextLine, len([]rune(fullTextLine)))
	}
}
