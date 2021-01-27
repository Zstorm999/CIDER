package editor

import (
	"fyne.io/fyne/widget"
	"github.com/Zstorm999/cider/explorer"
)

type Editor struct {
	Container *widget.TabContainer
}

func New() *Editor {

	e := &Editor{}

	e.Container = widget.NewTabContainer()

	return e

}

func (e *Editor) AddTab(file string) {

	newTab := widget.NewTabItem(explorer.ParseFileName(file), newEditorTab(file))

	e.Container.Append(newTab)
	e.Container.SelectTab(newTab)

}
