package explorer

import (
	"fmt"
	"os"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/container"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

type Explorer struct {
	path string

	folderName   *widget.Label
	folderSelect *widget.Button

	completion *widget.ProgressBar

	Container *fyne.Container
}

func (e *Explorer) UpdateTree(newPath string) {

	if newPath != e.path {

		e.completion.Show()

		e.path = newPath

		trimmedPath := strings.TrimPrefix(newPath, "file://")

		e.Container.Objects[0] = createFilesTree(trimmedPath, e.completion)

		e.completion.SetValue(1)

		e.folderName.SetText(parseFileName(newPath))

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

	e.folderName = widget.NewLabel(parseFileName(path))

	selectWindow := dialog.NewFolderOpen(
		func(file fyne.ListableURI, err error) {
			if file != nil {
				e.UpdateTree(file.String())
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
