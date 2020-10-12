package model

import "time"

// TimeModel data tipe to store time data
type TimeModel struct {
	Hour   uint8
	Minute uint8
	Second uint8
}

// NewTime TimeModel constructor
func NewTime() *TimeModel {
	t := &TimeModel{
		Hour:   0,
		Minute: 0,
		Second: 0,
	}

	t.Update()

	return t
}

// Update set time up to date
func (t *TimeModel) Update() {
	hour, minute, second := time.Now().Local().Clock()
	t.Hour = uint8(hour)
	t.Minute = uint8(minute)
	t.Second = uint8(second)
}
