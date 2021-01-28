package main

import (
	"log"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("ProgressBar Widget")

	progress := widget.NewProgressBar()
	infinite := widget.NewProgressBarInfinite()

	go func() {
		for i := 0.0; i <= 1.0; i += 0.1 {
			time.Sleep(time.Second * 1)
			progress.SetValue(i)
			log.Println(i)
		}
	}()

	myWindow.SetContent(container.NewVBox(progress, infinite))
	myWindow.ShowAndRun()
}
