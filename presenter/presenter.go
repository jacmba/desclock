package presenter

import (
	"github.com/jacmba/desclock/model"
	"github.com/jacmba/desclock/view"
)

// Presenter abstract data type for presentation layer
type Presenter interface {
	NewPresenter(m *interface{}) *Presenter
	InjectView(v *view.View)
	Update()
}

// PImpl implementation structure type for presentation layer
type PImpl struct {
	Model *model.TimeModel
	View  *view.View
}

// NewPresenter presenter constructor
func NewPresenter(m *model.TimeModel) *PImpl {
	p := &PImpl{
		Model: m,
		View:  nil,
	}

	return p
}

// Update update presenter
func (*PImpl) Update() {}
