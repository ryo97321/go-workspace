package main

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("KeyEvent Demo")

	text := canvas.NewText("", color.Black)
	text.Alignment = fyne.TextAlignCenter

	myWindow.SetContent(text)
	myWindow.Resize(fyne.NewSize(300, 300))

	myWindow.Canvas().SetOnTypedKey(func(ke *fyne.KeyEvent) {
		text.Refresh()
		if ke.Name == "Space" {
			text.Text = text.Text + " "
		} else {
			text.Text = text.Text + strings.ToLower(string(ke.Name))
		}
	})

	myWindow.ShowAndRun()
}
