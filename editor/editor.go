package editor

import "fyne.io/fyne/widget"

type editorTab struct {
	widget.BaseWidget

	content *widget.TextGrid
	focused bool
}

func newEditorTab() *editorTab {
	e := &editorTab{}
	//e.ExtendBaseWidget(e)

	e.content = widget.NewTextGrid()

	return e
}
