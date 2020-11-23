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

	var a, b, c, d int

	scanner.Scan()
	s := scanner.Text()
	params := strings.Split(s, " ")
	a, _ = strconv.Atoi(params[0])
	b, _ = strconv.Atoi(params[1])

	scanner.Scan()
	s = scanner.Text()
	params = strings.Split(s, " ")
	c, _ = strconv.Atoi(params[0])
	d, _ = strconv.Atoi(params[1])

	ans := a*d - b*c
	fmt.Println(ans)
}
