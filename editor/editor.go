package editor

import (
	"fyne.io/fyne/v2/widget"
	"github.com/Zstorm999/cider/explorer"
)

type editorTab struct {
	widget.BaseWidget

	content *widget.TextGrid
	focused bool
}

func newEditorTab(file string) *editorTab {
	e := &editorTab{}
	e.ExtendBaseWidget(e)

	e.content = widget.NewTextGrid()

	text := explorer.GetFileContent(file)

	if text == "" {
		text = "File could not be opened"
	}

	e.content.SetText(text)

	return e
}
