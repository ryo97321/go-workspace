package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"os"
)

func main() {
	// Open JPEG
	file, err := os.Open("guitar.jpg")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Decode jpeg
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}

	// Bound, RGBA
	imgBound := img.Bounds()
	out := image.NewRGBA(imgBound)

	// Main Process
	rect := out.Rect // RGBA -> Rectangle
	for y := rect.Min.Y; y < rect.Max.Y; y++ {
		for x := rect.Min.X; x < rect.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			r, g, b, a = r>>8, g>>8, b>>8, a>>8
			out.Set(x, y, color.RGBA{
				uint8(255 - r),
				uint8(255 - g),
				uint8(255 - b),
				255})
		}
	}

	file_out, err := os.Create("test.jpg")
	defer file_out.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Encode jpeg
	if err := jpeg.Encode(file_out, out, nil); err != nil {
		log.Fatal(err)
	}
}
