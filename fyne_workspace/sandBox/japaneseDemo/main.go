package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("こんにちは")

	hello := widget.NewLabel("こんにちは Fyne!")
	w.SetContent(widget.NewVBox(
		hello,
		widget.NewButton("こんにちは!", func() {
			hello.SetText("ようこそ :)")
		}),
	))

	w.ShowAndRun()
}
