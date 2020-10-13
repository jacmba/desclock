package main

import (
	"time"

	"github.com/jacmba/desclock/model"
	"github.com/jacmba/desclock/view"
)

func main() {
	tm := model.NewTime()
	v := view.NewView(tm)

	go run(v)

	v.Init()
}

func run(v *view.View) {
	v.Update()
	for {
		v.Update()
		time.Sleep(100 * time.Millisecond)
	}
}
