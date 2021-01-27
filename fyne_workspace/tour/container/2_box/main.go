package main

import (
	"fmt"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func sayHello() {
	fmt.Println("Hello!!")
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Entry Widget")

	content := container.NewVBox(
		widget.NewLabel("The top row of the VBox"),
		container.NewHBox(
			widget.NewLabel("Label 1"),
			widget.NewLabel("Label 2"),
		),
	)

	content.Add(widget.NewButton("Say Hello", sayHello))

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
