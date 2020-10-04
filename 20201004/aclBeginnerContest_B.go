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
	s := scanner.Text()

	params := strings.Split(s, " ")

	a, _ := strconv.Atoi(params[0])
	b, _ := strconv.Atoi(params[1])
	c, _ := strconv.Atoi(params[2])
	d, _ := strconv.Atoi(params[3])

	if b < c || d < a {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
