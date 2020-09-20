package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	repdigitCount := 0

	for i := 0; i < n; i++ {
		scanner.Scan()
		dices := strings.Split(scanner.Text(), " ")

		dice1 := dices[0]
		dice2 := dices[1]

		if dice1 == dice2 {
			repdigitCount++
		} else if dice1 != dice2 {
			repdigitCount = 0
		}

		if repdigitCount == 3 {
			break
		}
	}

	if repdigitCount == 3 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
