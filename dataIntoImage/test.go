package main

import (
	"fmt"
	"strconv"
)

func main() {
	// "01100111" -> 103
	foo, _ := strconv.ParseInt("01100111", 2, 0)
	fmt.Println(foo)
	fmt.Printf("%s\n", string(foo))

	testStr := "01100111"
	fmt.Println(testStr)

	lastBitChangedTestStr := testStr[:7] + "0"
	fmt.Println(lastBitChangedTestStr)
	fmt.Println(string(lastBitChangedTestStr[0]))
}
