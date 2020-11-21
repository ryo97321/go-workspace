package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func relu(x int) int {
	var ret int

	if x >= 0 {
		ret = x
	} else {
		ret = 0
	}

	return ret
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())

	fmt.Println(relu(x))
}
