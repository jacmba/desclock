package view

import "github.com/jacmba/desclock/presenter"

// View abstract data type to handle view stuff
type View interface {
	NewVieW(p *presenter.Presenter) *View
	Update(data *interface{})
}

// VImpl view implemented data type to handle view stuff
type VImpl struct {
	Presenter *presenter.Presenter
}
