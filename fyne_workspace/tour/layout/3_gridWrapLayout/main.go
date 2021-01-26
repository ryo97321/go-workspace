package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Grid Wrap Layout")

	text1 := canvas.NewText("1", color.Black)
	text2 := canvas.NewText("2", color.Black)
	text3 := canvas.NewText("3", color.Black)
	grid := container.New(layout.NewGridWrapLayout(fyne.NewSize(50, 50)), text1, text2, text3)

	w.SetContent(grid)
	w.ShowAndRun()
}
