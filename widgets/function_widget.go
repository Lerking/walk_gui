package widgets

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func intWidget() {
	var widget *walk.HBox
	var spinner *walk.NumberEdit
	var spinnerval int

	IntWidget{
		assignTo: &widget,
		Layout:  HBox{},
		Children: []Widget{
			Label{
				Text: "Integer",
			},
			NumberEdit{
				AssignTo: &spinner,
				Decimals: 0,
				SpinButtonsVisible: true,
				Value: &spinnerval,
			},
		},
	}
	return widget, spinner, spinnerval
}