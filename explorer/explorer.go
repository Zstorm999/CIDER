package explorer

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Explorer struct {
	path string

	folderName   *widget.Label
	folderSelect *widget.Button

	completion *widget.ProgressBar

	Container *fyne.Container

	lock sync.Mutex
}

func (e *Explorer) UpdateTree(newPath string) {

	e.lock.Lock()
	defer e.lock.Unlock()

	if newPath != e.path {

		e.completion.Show()

		e.path = newPath

		trimmedPath := strings.TrimPrefix(newPath, "file://")

		e.Container.Objects[0] = createFilesTree(trimmedPath, e.completion)

		e.completion.SetValue(1)

		e.folderName.SetText(ParseFileName(newPath))

		e.completion.Hide()
	}

}

//New : returns a fyne.Container containing a working file explorer
func New(win fyne.Window) *Explorer {

	//default application path
	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		path = ""
	}

	e := &Explorer{}

	e.path = path
	e.completion = widget.NewProgressBar()
	e.completion.Hide()

	treeFiles := createFilesTree(path, e.completion)

	e.folderName = widget.NewLabel(ParseFileName(path))

	selectWindow := dialog.NewFolderOpen(
		func(file fyne.ListableURI, err error) {
			if file != nil {
				go e.UpdateTree(file.String())
			}
		}, win)

	e.folderSelect = widget.NewButtonWithIcon("", theme.FolderIcon(),
		func() {
			selectWindow.Show()
		})

	topBar := container.NewBorder(nil, nil, e.folderName, e.folderSelect)

	e.Container = container.NewBorder(topBar, e.completion, nil, nil, treeFiles)

	return e

}
