package main

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

type skinnedEntry struct {
	widget.Entry
}

func newSkinnedEntry() *skinnedEntry {
	skinnedEntry := &skinnedEntry{}
	skinnedEntry.ExtendBaseWidget(skinnedEntry)
	return skinnedEntry
}

func (skinnedEntry *skinnedEntry) CreateRenderer() fyne.WidgetRenderer {
	skinnedEntry.ExtendBaseWidget(skinnedEntry)
	entryRenderer := skinnedEntry.Entry.CreateRenderer()
	return &skinnedEntryRenderer{
		skinnedEntry:  skinnedEntry,
		entryRenderer: entryRenderer,
	}
}

type skinnedEntryRenderer struct {
	skinnedEntry  *skinnedEntry
	entryRenderer fyne.WidgetRenderer
}

func (renderer *skinnedEntryRenderer) Refresh() {
	renderer.entryRenderer.Refresh()
}

func (renderer *skinnedEntryRenderer) MinSize() fyne.Size {
	return renderer.entryRenderer.MinSize()
}

func (renderer *skinnedEntryRenderer) Layout(size fyne.Size) {
	renderer.entryRenderer.Layout(size)
}

func (renderer *skinnedEntryRenderer) Objects() []fyne.CanvasObject {
	return renderer.entryRenderer.Objects()
}

func (renderer *skinnedEntryRenderer) Destroy() {
	renderer.entryRenderer.Destroy()
}

func (renderer *skinnedEntryRenderer) BackgroundColor() color.Color {
	return color.NRGBA{R: 0x7c, G: 0x7c, B: 0x7c, A: 0xff}
}

func main() {
	a := app.New()
	w := a.NewWindow("Skinned Entry")
	w.SetContent(newSkinnedEntry())
	w.ShowAndRun()
}
