package main

import "fmt"

func main() {
	blocks := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	for i := 0; i < len(blocks); i++ {
		for j := 0; j < len(blocks[i]); j++ {
			fmt.Printf("%d ", blocks[i][j])
		}
		fmt.Println()
	}
}
