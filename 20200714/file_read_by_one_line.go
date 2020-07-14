package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "file.txt"

	// File Open
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Make Scanner
	scanner := bufio.NewScanner(f)

	// Read by one-line
	// Scan advances the scanner to the next token
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
