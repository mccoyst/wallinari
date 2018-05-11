// © 2018 Steve McCoy under the MIT license. See LICENSE for details.

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"math/rand"
	"strconv"
)

func main() {
	xh, xw := 2436, 1125

	if len(os.Args) > 1 {
		seed, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		rand.Seed(int64(seed))
	}

	out := image.NewNRGBA(image.Rect(0, 0, xw, xh))

	divideAndContour(out, panels)

	file, err := os.Create("hot.png")
	if err != nil {
		panic(err)
	}
	err = png.Encode(file, out)
	if err != nil {
		panic(err)
	}
}

func divideAndContour(m *image.NRGBA, fill func(*image.NRGBA)) {
	q := 128
	for row := 0; row < m.Bounds().Max.Y; row += q {
	for col := 0; col < m.Bounds().Max.X; col += q {
		s := m.SubImage(image.Rect(col, row, col+q, row+q)).(*image.NRGBA)
		fill(s)
	}
	}
}

func panels(m *image.NRGBA) {
	offset := 32
	max := 256 - offset
	c := color.NRGBA{
		R: uint8(offset + rand.Intn(max)),
		G: uint8(offset + rand.Intn(max)),
		B: uint8(offset + rand.Intn(max)),
		A: 255,
	}

	u := image.NewUniform(c)
	draw.Draw(m, m.Bounds(), u, image.ZP, draw.Src)
}
