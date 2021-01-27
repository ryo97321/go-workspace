package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Border Layout")

	top := canvas.NewText("top bar", color.Black)
	left := canvas.NewText("left", color.Black)
	middle := canvas.NewText("content", color.Black)

	content := container.New(layout.NewBorderLayout(top, nil, left, nil), top, left, middle)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
