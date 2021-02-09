package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("こんにちは")

	hello := widget.NewLabel("こんにちは Fyne!")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("こんにちは!", func() {
			hello.SetText("ようこそ :)")
		}),
	))

	w.ShowAndRun()

	// Run command
	// FYNE_FONT=C:\\Windows\\Fonts\\Meiryo.ttc go run main.go
}
