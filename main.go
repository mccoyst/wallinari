// © 2018 Steve McCoy under the MIT license. See LICENSE for details.

package main

import (
	"image"
	"image/color"
	"image/color/palette"
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
	u := image.NewUniform(color.White)
	draw.Draw(out, out.Bounds(), u, image.ZP, draw.Src)

	divideAndContour(out, glenda)

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

func dirt(m *image.NRGBA) {
	bounds := m.Bounds()
	dx, dy := bounds.Dx(), m.Bounds().Dy()
	ndots :=  dx * dy / 16
	for i := 0; i < ndots; i++ {
		x := rand.Intn(dx)
		y := rand.Intn(dy)
		m.Set(bounds.Min.X + x, bounds.Min.Y + y, color.Black)
		m.Set(bounds.Min.X + x - 1, bounds.Min.Y + y, color.Black)
		m.Set(bounds.Min.X + x + 1, bounds.Min.Y + y, color.Black)
		m.Set(bounds.Min.X + x, bounds.Min.Y + y - 1, color.Black)
		m.Set(bounds.Min.X + x, bounds.Min.Y + y + 1, color.Black)
	}
}

func kitchen(m *image.NRGBA) {
	if rand.Intn(2) == 1 {
		u := image.NewUniform(color.Black)
		draw.Draw(m, m.Bounds(), u, image.ZP, draw.Src)
	}
}

func glenda(m *image.NRGBA) {
	c := palette.Plan9[rand.Intn(len(palette.Plan9))]
	u := image.NewUniform(c)
	draw.Draw(m, m.Bounds(), u, image.ZP, draw.Src)
}

func sticks(m *image.NRGBA) {
	 // Bresenham
}
