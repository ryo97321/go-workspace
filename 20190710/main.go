package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func grad(img *image.RGBA) {
	rect := img.Rect
	for h := rect.Min.Y; h < rect.Max.Y; h++ {
		for v := rect.Min.X; v < rect.Max.X; v++ {
			img.Set(v, h, color.RGBA{
				uint8(255 * h / rect.Max.Y),
				uint8(255 * v / rect.Max.X),
				uint8(255 * v * h / (rect.Max.X * rect.Max.Y)),
				0})
		}
	}
}

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	grad(img)

	file, err := os.Create("test.jpg")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := jpeg.Encode(file, img, &jpeg.Options{100}); err != nil {
		log.Fatal(err)
	}
}
