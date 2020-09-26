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
	n, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	var a int64
	var ans int64 = 0
	for a = 1; a <= n; a++ {
		ans += (n - 1) / a
	}

	fmt.Println(ans)
}
