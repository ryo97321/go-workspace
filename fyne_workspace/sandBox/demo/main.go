package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("ButtonDemo")

	text := canvas.NewText("button", color.White)
	w.SetContent(text)

	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
