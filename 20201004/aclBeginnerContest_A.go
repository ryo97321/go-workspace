package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	k, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < k; i++ {
		fmt.Print("ACL")
	}
	fmt.Println()
}
