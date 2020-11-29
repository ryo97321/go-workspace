package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	s := scanner.Text()
	params := strings.Split(s, " ")
	h, _ := strconv.Atoi(params[0])
	w, _ := strconv.Atoi(params[1])

	grid := make([][]int, h)
	for i := 0; i < h; i++ {
		scanner.Scan()
		s = scanner.Text()
		params = strings.Split(s, " ")
		row := make([]int, 0)
		for j := 0; j < w; j++ {
			cost, _ := strconv.Atoi(params[j])
			row = append(row, cost)
		}
		grid[i] = row
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, -1, 0, 1}
	R, U, L, D := 0, 1, 2, 3

	totalCost := 0
	x, y := 0, 0 // 現在のx座標, y座標

	// 初期位置のコストを加算
	totalCost += grid[y][x]

	// 右
	x += dx[R]
	y += dy[R]
	totalCost += grid[y][x]

	// 下
	x += dx[D]
	y += dy[D]
	totalCost += grid[y][x]

	// 右
	x += dx[R]
	y += dy[R]
	totalCost += grid[y][x]

	// 上
	x += dx[U]
	y += dy[U]
	totalCost += grid[y][x]

	// 左
	x += dx[L]
	y += dy[L]
	totalCost += grid[y][x]

	fmt.Println(totalCost)
}
