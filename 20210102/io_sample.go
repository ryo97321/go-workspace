package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "読み込みに失敗しました：", err)
	// }

	err := filepath.Walk("dir",
		func(path string, info os.FileInfo, err error) error {
			if filepath.Ext(path) == ".go" {
				fmt.Println(path)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
}
