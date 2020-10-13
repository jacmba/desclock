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
	analogFontSize      = 18.0
	pi2                 = 2.0 * 3.141592
	piMed               = 3.141592 / 2.0
	pi                  = 3.141592
	hourRadiusPercent   = 0.50
	minuteRadiusPercent = 0.70
	secondRadiusPercent = 0.85
	hourWidth           = 5.0
	minuteWidth         = 3.0
	secondWith          = 1.0
)

// Draw analog clock
func renderAnalog(t *model.TimeModel, w, h int) image.Image {
	ctx := gg.NewContext(w, h)
	if err := ctx.LoadFontFace("./res/font/display.ttf", analogFontSize); err != nil {
		panic(err)
	}
	ctx.SetColor(color.White)

	// Get center and radius
	cx := float64(w) / 2.0
	cy := float64(h) / 2.0
	r := math.Min(cx, cy) * 0.9

	nStep := pi2 / 12.0 // Numbers separation in angles
	fStep := pi2 / 60.0 // Angle separation for minutes and seconds

	// Draw clock numbers
	for n := 1; n <= 12; n++ {
		alpha := float64(n)*nStep - piMed // Get angle for current number
		x := r*math.Cos(alpha) + cx       // Horizontal coordinate
		y := r*math.Sin(alpha) + cy       // Vertical coordinate

		sn := fmt.Sprint(n)
		ctx.DrawString(sn, x, y)
	}

	// Draw seconds
	rs := r * secondRadiusPercent
	sAlpha := float64(t.Second)*fStep - piMed
	sx := rs*math.Cos(sAlpha) + cx
	sy := rs*math.Sin(sAlpha) + cy
	ctx.SetLineWidth(secondWith)
	ctx.DrawLine(cx, cy, sx, sy)
	ctx.Stroke()

	// Draw minutes
	rm := r * minuteRadiusPercent
	mAlpha := float64(t.Minute)*fStep - piMed
	mx := rm*math.Cos(mAlpha) + cx
	my := rm*math.Sin(mAlpha) + cy
	ctx.SetLineWidth(minuteWidth)
	ctx.DrawLine(cx, cy, mx, my)
	ctx.Stroke()

	// Draw hours
	rh := r * hourRadiusPercent
	hour := float64(t.Hour)
	if hour > 12 {
		hour -= 12
	}
	hAlpha := hour*nStep - piMed
	hx := rh*math.Cos(hAlpha) + cx
	hy := rh*math.Sin(hAlpha) + cy
	ctx.SetLineWidth(hourWidth)
	ctx.DrawLine(cx, cy, hx, hy)
	ctx.Stroke()

	return ctx.Image()
}
