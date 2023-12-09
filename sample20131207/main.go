package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0
	for i := 0; i < 5; i++ {
		fmt.Print("input number: ")
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())

		sum += n
	}
	fmt.Printf("sum = %d\n", sum)

	/* No30から再開 */
	/* https://www.cc.kyoto-su.ac.jp/~mmina/bp1/hundredKnocksPrimary.html */
}
