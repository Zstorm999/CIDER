package editor

import (
	"fyne.io/fyne/v2/container"
	"github.com/Zstorm999/cider/explorer"
)

type Editor struct {
	Container *container.AppTabs
}

func New() *Editor {

	e := &Editor{}

	e.Container = container.NewAppTabs()

	return e

}

func (e *Editor) AddTab(file string) {

	newTab := container.NewTabItem(explorer.ParseFileName(file), newEditorTab(file))

	e.Container.Append(newTab)
	e.Container.SelectTab(newTab)

}
