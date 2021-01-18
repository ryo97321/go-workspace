package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(widget.NewHBox(
		hello,
		widget.NewButton("Yes!", func() {
			hello.SetText("Hello! Yes!")
		}),
		widget.NewButton("No!", func() {
			hello.SetText("Hello! No!")
		}),
	))

	w.ShowAndRun()
}
