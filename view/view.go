package view

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"github.com/jacmba/desclock/model"
)

const (
	// DigitalView value to show digital clock
	DigitalView = 1

	// AnalogView value to show analog clock
	AnalogView = 2

	width = 400

	height = 400
)

// View data structure to process view
type View struct {
	mode   uint8
	model  *model.TimeModel
	aPP    fyne.App
	window fyne.Window
}

// NewView constructor of View data type
func NewView(m *model.TimeModel) *View {
	a := app.New()
	w := a.NewWindow("Desclock")
	v := View{
		mode:   AnalogView,
		model:  m,
		aPP:    a,
		window: w,
	}

	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(width, height))
	w.CenterOnScreen()
	return &v
}

// Init initalize the window system
func (v *View) Init() {
	v.window.ShowAndRun()
}

// Update refresh view
func (v *View) Update() {
	v.model.Update()
	if v.mode == DigitalView {
		v.renderDigital()
	} else {
		v.renderAnalog()
	}
}

func (v *View) renderDigital() {
	h := strTime(v.model.Hour)
	m := strTime(v.model.Minute)
	s := strTime(v.model.Second)
	t := fmt.Sprintf("%s:%s:%s", h, m, s)
	img := renderDigital(t, v.window.Canvas().Size().Width, v.window.Canvas().Size().Height)
	raster := canvas.NewRasterFromImage(img)
	v.window.SetContent(raster)
}

func (v *View) renderAnalog() {
	img := renderAnalog(v.model, v.window.Canvas().Size().Width, v.window.Canvas().Size().Height)
	raster := canvas.NewRasterFromImage(img)
	v.window.SetContent(raster)
}

func strTime(t uint8) string {
	s := fmt.Sprintf("%d", t)
	if t < 10 {
		s = "0" + s
	}
	return s
}
