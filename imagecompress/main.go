package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

// converts og_search.png with various methods
//
// og_search.png (26.8kB) : source file
// 0. og_search_resize.png (9.63kB) : re-sized (400 x 200 px) file, others are based on re-sized file.

// 1. og_search_custom.png (704B) : using two colors, white and blue, set color to closer one.
// 2. og_search_grey.png (3.97kB) : gray-scaling the image
// 3. og_search_monotone.png (1.35kB) : customized logic, when average of sum of r, g and b are bigger than threshold, make it white.
// 4. og_search_plan9.png (2.49kB) : using color palettes of plan9 colors
// 5. og_serach_websafe.png (2.27kB) : using color palettes of web-safe colors
func main() {
	fileName := "og_search.png"

	read, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer read.Close()

	src, err := png.Decode(read)
	if err != nil {
		fmt.Println(err)
	}

	src = resize.Resize(400, 200, src, resize.Lanczos3)

	outFile := "og_search_resize.png"
	write, err := os.Create(outFile)
	if err != nil {
		fmt.Println(err)
	}
	defer write.Close()
	err = png.Encode(write, src)
	if err != nil {
		fmt.Println(err)
	}

	err = toGrayScale(src)
	err = toPalette(src, "websafe")
	err = toPalette(src, "plan9")
	err = toPalette(src, "custom")
	err = removeWhite(src)
	if err != nil {
		fmt.Println(err)
	}

}

func removeWhite(src image.Image) error {
	dest := image.NewRGBA(src.Bounds())
	c := color.RGBA{99, 146, 249, 255}

	for x := src.Bounds().Min.X; x <= src.Bounds().Max.X; x++ {
		for y := src.Bounds().Min.Y; y <= src.Bounds().Max.Y; y++ {
			r, g, b, _ := src.At(x, y).RGBA()
			th := ((r + g + b) / 3) >> 8
			if th >= 190 {
				dest.Set(x, y, color.White)
			} else {
				dest.Set(x, y, c)
			}
		}
	}
	outFile := "og_search_monotone.png"
	write, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer write.Close()
	err = png.Encode(write, dest)
	if err != nil {
		return err
	}

	return nil
}

func toPalette(src image.Image, t string) error {

	custom := color.Palette{
		color.RGBA{99, 146, 249, 255},
		color.RGBA{255, 255, 255, 255},
	}

	var p color.Palette
	switch t {
	case "plan9":
		p = palette.Plan9
	case "websafe":
		p = palette.WebSafe
	default:
		p = custom
	}

	dest := image.NewPaletted(src.Bounds(), p)
	for x := src.Bounds().Min.X; x <= src.Bounds().Max.X; x++ {
		for y := src.Bounds().Min.Y; y <= src.Bounds().Max.Y; y++ {
			dest.Set(x, y, src.At(x, y))
		}
	}
	outFile := fmt.Sprintf("og_search_%v.png", t)
	write, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer write.Close()
	err = png.Encode(write, dest)
	if err != nil {
		return err
	}
	return nil
}

func toGrayScale(src image.Image) error {
	dest := image.NewGray(src.Bounds())
	for x := src.Bounds().Min.X; x <= src.Bounds().Max.X; x++ {
		for y := src.Bounds().Min.Y; y <= src.Bounds().Max.Y; y++ {
			dest.Set(x, y, src.At(x, y))
		}
	}
	outFile := "og_search_grey.png"
	write, err := os.Create(outFile)
	if err != nil {
		return err
	}
	defer write.Close()
	err = png.Encode(write, dest)
	if err != nil {
		return err
	}
	return nil
}
