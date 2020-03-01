package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var img = image.NewRGBA(image.Rect(0, 0, 100, 100))
var col color.Color

func HLine(x1, y, x2 int) {
	for ; x1 <= x2; x1++ {
		img.Set(x1, y, col)
	}
}

func main() {
	var i int
	col = color.RGBA{255, 0, 0, 255} // Red
	for i = 0; i < 100; i++ {
		if i%2 != 0 {
			HLine(0, i, 100)
		}
	}

	f, err := os.Create("draw.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}
