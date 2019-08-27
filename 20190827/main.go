package main

import (
	"fmt"
	"math/rand"

	"github.com/goml/gobrain"
)

func main() {
	// set the random seed to 0
	rand.Seed(0)

	// xorの入力とその入力に対する答え（教師データ）
	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	// モデル初期化
	ff := &gobrain.FeedForward{}

	// ネットワークを構成
	// 引数は左から順に、入力数、中間層の次元数、答えの数を設定する
	ff.Init(2, 2, 1)

	// 学習させるメソッド
	// 引数は左から順に学習させるパターン、学習回数、学習率、運動量係数、最後のエラーを出力するかどうか
	ff.Train(patterns, 1000, 0.6, 0.4, false)

	inputs := [][]float64{
		{1, 1},
		{1, 0},
		{0, 1},
		{0, 0},
	}
	for _, input := range inputs {
		result := ff.Update(input)
		fmt.Println(input[0], ",", input[1], ":", result[0])
	}
}
