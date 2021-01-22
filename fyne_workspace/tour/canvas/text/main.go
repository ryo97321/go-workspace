package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("text")

	text := canvas.NewText("Text Object", color.Black)
	text.Alignment = fyne.TextAlignTrailing
	text.TextStyle = fyne.TextStyle{Italic: true}
	w.SetContent(text)

	w.ShowAndRun()
}
