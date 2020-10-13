package view

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/fogleman/gg"
	"github.com/jacmba/desclock/model"
)

const (
	analogFontSize = 18.0
	pi2            = 2.0 * 3.141592
	piMed          = 3.141592 / 2.0
	pi             = 3.141592
)

// Draw analog clock
func renderAnalog(t *model.TimeModel, w, h int) image.Image {
	ctx := gg.NewContext(w, h)
	if err := ctx.LoadFontFace("./res/font/display.ttf", analogFontSize); err != nil {
		panic(err)
	}
	ctx.SetColor(color.White)
	ctx.SetLineWidth(20)

	// Get center and radius
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := math.Min(cx, cy) * 0.9

	ctx.DrawLine(cx, cy, 400, 400)

	nStep := pi2 / 12.0 // Numbers separation in angles
	//fStep := pi2 / 60.0 // Angle separation for minutes and seconds

	// Draw clock numbers
	for n := 1; n <= 12; n++ {
		alpha := float64(n)*nStep - piMed // Get angle for current number
		x := r*math.Cos(alpha) + cx       // Horizontal coordinate
		y := r*math.Sin(alpha) + cy       // Vertical coordinate

		sn := fmt.Sprint(n)
		ctx.DrawString(sn, x, y)
	}

	// Draw seconds
	//rs := r * secondRadiusPercent
	//sAlpha := float64(t.Second)*fStep - piMed
	//sx := r*math.Cos(sAlpha) + cx
	//sy := r*math.Sin(sAlpha) + cy

	return ctx.Image()
}
