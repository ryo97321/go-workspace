package main

import (
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("BMI")

	heightLabel := canvas.NewText("Height (cm)", color.Black)
	heightInput := widget.NewEntry()
	heightContent := container.New(layout.NewHBoxLayout(), heightLabel, heightInput)

	weightLabel := canvas.NewText("Weight (kg)", color.Black)
	weightInput := widget.NewEntry()
	weightContent := container.New(layout.NewHBoxLayout(), weightLabel, weightInput)

	bmiLabel := canvas.NewText("BMI : ", color.Black)
	bmiResultLabel := canvas.NewText("", color.Black)
	bmiContent := container.New(layout.NewHBoxLayout(), bmiLabel, bmiResultLabel)

	calcButton := widget.NewButton("Calc", func() {
		bmiResultLabel.Refresh()

		heightText := heightInput.Text
		weightText := weightInput.Text
		height, _ := strconv.ParseFloat(heightText, 64)
		weight, _ := strconv.ParseFloat(weightText, 64)
		height = height / 100.0 // cm -> m
		bmi := weight / (height * height)

		bmiString := strconv.FormatFloat(bmi, 'f', 4, 64)
		bmiResultLabel.Text = bmiString
	})

	content := container.New(layout.NewVBoxLayout(), heightContent, weightContent, bmiContent, calcButton)

	myWindow.Resize(fyne.NewSize(300, 100))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
