package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var arr [5]int
	var n int
	for i := 0; i < 5; i++ {
		fmt.Print("input number: ")
		scanner.Scan()
		n, _ = strconv.Atoi(scanner.Text())
		arr[i] = n
	}

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}

	/* No20から再開 */
	/* http://www.cc.kyoto-su.ac.jp/~mmina/bp1/hundredKnocksPrimary.html */
}
