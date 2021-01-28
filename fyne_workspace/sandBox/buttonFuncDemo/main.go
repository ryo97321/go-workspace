package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type myButton struct {
	widget.Button
}

// Create Extended widget.Button
func newMyButton(label string) *myButton {
	button := &myButton{}
	button.ExtendBaseWidget(button)
	button.SetText(label)

	return button
}

func (mb *myButton) MouseIn(_ *desktop.MouseEvent) {
	log.Println("In")
}

func (mb *myButton) MouseOut() {
	log.Println("Out")
}

func (mb *myButton) Tapped(pe *fyne.PointEvent) {
	fmt.Printf("(%f, %f)\n", pe.Position.X, pe.Position.Y)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Button Func Demo")

	text := canvas.NewText("Button Func Demo", color.Black)
	text.Alignment = fyne.TextAlignCenter

	button := newMyButton("myButton")

	content := container.New(layout.NewVBoxLayout(), text, button)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
