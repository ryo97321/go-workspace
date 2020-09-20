package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	sEnd := string(s[len(s)-1])

	result := ""
	if sEnd == "s" {
		result = s + "es"
	} else {
		result = s + "s"
	}

	fmt.Println(result)
}
