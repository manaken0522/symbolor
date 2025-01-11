package main

import (
	"bytes"
	"flag"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

func main() {
	flag.Parse()
	file_path := flag.Arg(0)

	var (
		img         image.Image
		red_total   int
		green_total int
		blue_total  int
	)

	filebyte, error := os.ReadFile(file_path)
	if error != nil {
		panic(error)
	}

	reader := bytes.NewReader(filebyte)

	extension := filepath.Ext(file_path)
	switch extension {
	case ".png":
		img, _ = png.Decode(reader)
	case ".jpeg", ".jpg":
		img, _ = jpeg.Decode(reader)
	}
	if error != nil {
		panic(error)
	}

	x_max := img.Bounds().Max.X
	y_max := img.Bounds().Max.Y

	for x := 0; x != x_max; x++ {
		for y := 0; y != y_max; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			red_total += int(r >> 8)
			green_total += int(g >> 8)
			blue_total += int(b >> 8)
		}
	}

	red := red_total / (x_max * y_max)
	green := green_total / (x_max * y_max)
	blue := blue_total / (x_max * y_max)
	color.RGB(red, green, blue).Printf("R:%d, G:%d, B:%d\n■■■■■■■■■■", red, green, blue)
}
