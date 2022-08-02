package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type main_struct struct {
	MainWidget     *walk.MainWindow
	MainListbox    *walk.ListBox
	MainScrollView *walk.ScrollView
}

func main() {
	var MainWin = &main_struct{}

	MainWindow{
		AssignTo: &MainWin.MainWidget,
		Title:    "PM rules",
		MinSize:  Size{800, 600},
		Size:     Size{800, 600},
		Layout:   HBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					ListBox{
						AssignTo: &MainWin.MainListbox,
						Model:    MainWin.MainListbox.Model,
						OnCurrentIndexChanged: func() {
							MainWin.MainScrollView.SetModel(MainWin.MainListbox.Model)
						},
					},
					ScrollView{
						AssignTo: &MainWin.MainScrollView,
						Layout:   VBox{},

			}
	}.Run()
}
