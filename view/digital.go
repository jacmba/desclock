package view

import (
	"image"
	"image/color"

	"github.com/fogleman/gg"
)

const (
	fontSize = 48.0
)

func renderDigital(s string, w, h int) image.Image {
	ctx := gg.NewContext(w, h)
	if err := ctx.LoadFontFace("./res/font/display.ttf", fontSize); err != nil {
		panic(err)
	}
	ctx.SetColor(color.White)
	x := float64(w) / 2.0
	y := float64(h) / 2.0
	ctx.DrawStringAnchored(s, x, y, 0.5, 0.5)
	return ctx.Image()
}
