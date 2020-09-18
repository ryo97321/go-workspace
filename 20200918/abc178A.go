package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	x, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	if x == 0 {
		x = 1
	} else if x == 1 {
		x = 0
	}

	fmt.Println(x)
}
