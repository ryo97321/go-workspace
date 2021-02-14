package main

import (
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	fullText := "This\nis\nsample\ntext\n"
	fullTextLines := strings.Split(fullText, "\n")
	nLine := len(fullTextLines)

	texts := make([]fyne.CanvasObject, nLine)
	for _, line := range fullTextLines {
		text := canvas.NewText(line, color.Black)
		text.Alignment = fyne.TextAlignCenter
		texts = append(texts, text)
	}

	app := app.New()
	w := app.NewWindow("Title")

	content := container.New(layout.NewVBoxLayout(), texts...)

	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
