package widgets

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func featuresWidget() {
	var FWidget *walk.ScrollView

	ScrollView{
		AssignTo: &FWidget,
		Children: []Widget{
			Label{
				Text: "Input",
			},
			Composite{
				Layout: VBox{},
				Children: []Widget{
					CheckBox{
						Text: "Unspecified",
					},
					CheckBox{
						Text: "Not",
					},
				},
			},
			ComboBox{
				Model: []string{"Unspecified", "Not"},
			},
		},
	}
}
