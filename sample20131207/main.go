package main

import (
	"fmt"
)

func main() {
	arr := [10]int{3, 7, 0, 8, 4, 1, 9, 6, 5, 2}

	for i := 0; i < 9; i++ {
		fmt.Println(arr[i] - arr[i+1])
	}

	/* 基礎プロI 100本ノック 中級編から再開 */
	/* https://www.cc.kyoto-su.ac.jp/~mmina/bp1/hundredKnocks.html */
}
