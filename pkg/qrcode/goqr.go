package qrcode

import (
	"errors"
	"image"
	"image/png"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/fatih/color"
)

//PrintCode prints a qrcode given string
func PrintCode(data string) {
	size := 880
	step := int(size / 22)

	qrCode, err := qr.Encode(data, qr.M, qr.Auto)
	if err != nil {
		panic(errors.New("Problem gnereating qrcode"))
	}
	qrCode, _ = barcode.Scale(qrCode, size, size)

	pixels, err := getPixels(qrCode)
	if err != nil {
		panic(errors.New("Promlom decodeing image"))
	}

	printCode(pixels, step)
}

func getPixels(img barcode.Barcode) ([][]bool, error) {
	image.RegisterFormat("png", "png", png.Decode, png.DecodeConfig)

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]bool
	for y := 0; y < height; y++ {
		var row []bool
		for x := 0; x < width; x++ {
			row = append(row, rgbaToBool(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

func rgbaToBool(r uint32, g uint32, b uint32, a uint32) bool {
	if int(r/257) == 0 &&
		int(g/257) == 0 &&
		int(b/257) == 0 &&
		int(a/257) == 255 {
		return true
	}
	return false
}

func printCode(pixels [][]bool, step int) {
	whiteBackground := color.New(color.FgBlack).Add(color.BgHiWhite)
	white := color.New(color.FgHiWhite)
	defer color.Unset()

	for y := 0; y < (len(pixels) - step); y += step * 2 {
		for x := 0; x < len(pixels[y]); x += step {
			if pixels[y][x] && pixels[y+step][x] { // full
				whiteBackground.Print("\u2588") //█

			} else if !pixels[y][x] && !pixels[y+step][x] { // empty
				whiteBackground.Print("\u2003") // \

			} else if pixels[y][x] && !pixels[y+step][x] { // top
				whiteBackground.Print("\u2580") //▀

			} else if !pixels[y][x] && pixels[y+step][x] { // bottom
				whiteBackground.Print("\u2584") //▄
			}
		}
		whiteBackground.Print("\u2003")
		white.Println()
	}
	color.Unset()

	for x := 0; x < len(pixels[0])+1; x += step {
		white.Print("\u2580")
	}
	white.Println()
}
