package main

import "fmt"

// fullText を30文字ごとに分割する
func splitFullTextEvery30Chars(runes []rune) []string {
	var ss []string

	s := ""
	for i, r := range runes {
		s += string(r)

		if (i+1)%30 == 0 { // 30文字になったらsをssに加え、sを空にする
			ss = append(ss, s)
			s = ""
		}
	}
	if s != "" { // s に文字列が残っていたらssに加える
		ss = append(ss, s)
	}

	return ss
}

func main() {
	fullText := "月曜日はMondayです。火曜日はTuesdayです。水曜日はWednesdayです。"
	runes := []rune(fullText) // string -> []rune

	fullTextLines := splitFullTextEvery30Chars(runes)
	fmt.Printf("全%d行, %d文字\n", len(fullTextLines), len(runes))
	for _, fullTextLine := range fullTextLines {
		fmt.Printf("%s : %d文字\n", fullTextLine, len([]rune(fullTextLine)))
	}
}
