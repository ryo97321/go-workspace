package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/axamon/gobrain/persist"
	"github.com/goml/gobrain"
)

func main() {
	rand.Seed(0)

	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	ff := &gobrain.FeedForward{}

	ff.Init(2, 2, 1)

	ff.Train(patterns, 1000, 0.6, 0.4, false)
	filename := "./ff.json"
	err := persist.Save(filename, ff)
	if err != nil {
		log.Println("Failed Save: ", err.Error())
	}

	ff2 := &gobrain.FeedForward{}
	err = persist.Load(filename, &ff2)
	if err != nil {
		log.Println("Failed Load: ", err.Error())
	}

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
