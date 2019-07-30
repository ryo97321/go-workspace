package main

import (
	"context"

	runewidth "github.com/mattn/go-runewidth"

	"github.com/eihigh/goban"
)

func main() {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
	goban.Main(app, view)
}

func app(_ context.Context, es goban.Events) error {
	goban.Show()
	es.ReadKey()
	return nil
}

func view() {
	goban.Screen().Enclose("hello").Prints("Hello World!\nPress any key to exit.")
}
