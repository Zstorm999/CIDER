package editor

import (
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

func New() *widget.TabContainer {

	textEditor1 := newEditorTab()
	textEditor2 := newEditorTab()

	return container.NewAppTabs(
		container.NewTabItem("Tab1", textEditor1),
		container.NewTabItem("Tab2", textEditor2))

}
