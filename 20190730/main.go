package main

import (
	"context"

	runewidth "github.com/mattn/go-runewidth"

	"github.com/eihigh/goban"
)

func main() {
	runewidth.DefaultCondition = &runewidth.Condition{EastAsianWidth: false}
	goban.Main(app)
}

func app(_ context.Context, es goban.Events) error {
	v := func() {
		b := goban.Screen()
		b.Puts("Press any key to open popup")
		b.Puts("Ctrl+C to exit")
	}
	goban.PushViewFunc(v)

	for {
		goban.Show()
		es.ReadKey()
		popup(es)
	}
}

func popup(es goban.Events) {
	v := func() {
		b := goban.NewBox(0, 0, 40, 5).Enclose("popup window")
		b.Prints("Press any key to popup")
	}

	goban.PushViewFunc(v)
	defer goban.PopView()

	goban.Show()
	es.ReadKey()
}
