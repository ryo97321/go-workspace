package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var n *bool
var nLine int = 1 // 行番号（通し番号）

func readFile(filePath string, opt bool) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if opt == true { // -n オプションが指定されているとき
		for scanner.Scan() {
			fmt.Printf("%d : %s\n", nLine, scanner.Text())
			nLine++
		}
	} else { // -n オプションが指定されていないとき
		for scanner.Scan() {
			fmt.Printf("%s\n", scanner.Text())
		}
	}

	return nil
}

func main() {
	n = flag.Bool("n", false, "行番号を表示するかしないか")
	flag.Parse()

	filePaths := flag.Args()
	if len(filePaths) != 0 {
		for _, filePath := range filePaths {
			err := readFile(filePath, *n)
			if err != nil {
				fmt.Println("err: ", err)
			}
		}
	}
}
