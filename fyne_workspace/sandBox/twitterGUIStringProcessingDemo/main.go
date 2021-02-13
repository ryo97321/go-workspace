package main

import (
	"fmt"
	"image/color"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func getTweetFullTextLines(fullText string) []string {
	fullTextLines := strings.Split(fullText, "\n")

	return fullTextLines
}

func main() {
	tweetFullText := "My\nname\nis\n***\n"

	tweetFullTextLines := getTweetFullTextLines(tweetFullText)
	for _, line := range tweetFullTextLines {
		fmt.Println(line)
	}

	app := app.New()
	w := app.NewWindow("Title")

	text1 := canvas.NewText("Text1", color.Black)
	text1.Alignment = fyne.TextAlignCenter
	text2 := canvas.NewText("Text2", color.Black)
	text2.Alignment = fyne.TextAlignCenter
	text3 := canvas.NewText("Text3", color.Black)
	text3.Alignment = fyne.TextAlignCenter

	content := container.New(layout.NewVBoxLayout(), text1, text2, text3)

	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 300))
	w.ShowAndRun()
}
