package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isFirst(winningNumber, number int) bool {
	if winningNumber == number {
		return true
	}

	return false
}

func isSecond(winningNumber, number int) bool {
	winningNumberStr := strconv.Itoa(winningNumber)
	numberStr := strconv.Itoa(number)

	winningNumberSlice := strings.Split(winningNumberStr, "")
	numberSlice := strings.Split(numberStr, "")

	result := true
	for i := len(numberSlice) - 1; i >= 2; i-- {
		if winningNumberSlice[i] != numberSlice[i] {
			result = false
			break
		}
	}

	return result
}

func isThird(winningNumber, number int) bool {
	winningNumberStr := strconv.Itoa(winningNumber)
	numberStr := strconv.Itoa(number)

	winningNumberSlice := strings.Split(winningNumberStr, "")
	numberSlice := strings.Split(numberStr, "")

	result := true
	for i := len(numberSlice) - 1; i >= 3; i-- {
		if winningNumberSlice[i] != numberSlice[i] {
			result = false
			break
		}
	}

	return result
}

func isAdjacent(winningNumber, number int) bool {
	if number == winningNumber-1 || number == winningNumber+1 {
		return true
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	winningNumber, _ := strconv.Atoi(s)

	scanner.Scan()
	s = scanner.Text()
	n, _ := strconv.Atoi(s)

	for i := 0; i < n; i++ {
		scanner.Scan()
		s = scanner.Text()
		number, _ := strconv.Atoi(s)

		msg := "blank"
		if isFirst(winningNumber, number) {
			msg = "first"
		} else if isAdjacent(winningNumber, number) {
			msg = "adjacent"
		} else if isSecond(winningNumber, number) {
			msg = "second"
		} else if isThird(winningNumber, number) {
			msg = "third"
		}

		fmt.Println(msg)
	}
}
