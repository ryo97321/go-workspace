package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TypingGame Demo")

	text := canvas.NewText("Typing", color.RGBA{0x80, 0x80, 0x80, 0xff})
	text.Alignment = fyne.TextAlignCenter
	text.TextSize = 64

	myWindow.SetContent(text)
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}
