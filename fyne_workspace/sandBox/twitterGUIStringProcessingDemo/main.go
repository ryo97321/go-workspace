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

	app := app.New()
	w := app.NewWindow("Title")
	w.Resize(fyne.NewSize(300, 300))

	texts := make([]fyne.CanvasObject, 0, nLine)
	for _, line := range fullTextLines {
		text := canvas.NewText(line, color.Black)
		text.Alignment = fyne.TextAlignCenter
		texts = append(texts, text)
	}

	content := container.New(layout.NewVBoxLayout(), texts...)

	w.SetContent(content)
	w.ShowAndRun()
}
