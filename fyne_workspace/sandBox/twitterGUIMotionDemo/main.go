package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func setNewText(text *canvas.Text, w fyne.Window) {
	count := 0
	for {
		time.Sleep(time.Second * 1)
		count++

		text.Refresh()
		if count%2 == 0 {
			text.Text = "Even"
		} else {
			text.Text = "Odd"
		}

		w.SetContent(text)
		w.Show()
	}
}

func main() {
	app := app.New()
	w := app.NewWindow("TwitterGUI Design Demo")

	text := canvas.NewText("Text", color.Black)
	text.Alignment = fyne.TextAlignCenter
	w.SetContent(text)
	w.Resize(fyne.NewSize(300, 300))

	go setNewText(text, w)

	w.ShowAndRun()
}
