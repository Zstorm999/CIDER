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

	Container *fyne.Container
}

func (e *Explorer) UpdateTree(newPath string) {
	if newPath != e.path {
		e.path = newPath

		trimmedPath := strings.TrimPrefix(newPath, "file://")
		e.Container.Objects[0] = createFilesTree(trimmedPath)

		sliced := strings.Split(newPath, "/")
		n := len(sliced)
		e.folderName.SetText(sliced[n-1])

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
	treeFiles := createFilesTree(path)

	sliced := strings.Split(path, "/")
	n := len(sliced)
	e.folderName = widget.NewLabel(sliced[n-1])

	selectWindow := dialog.NewFolderOpen(
		func(file fyne.ListableURI, err error) {
			e.UpdateTree(file.String())
		}, win)

	e.folderSelect = widget.NewButtonWithIcon("", theme.FolderIcon(),
		func() {
			selectWindow.Show()
		})

	topBar := container.NewBorder(nil, nil, e.folderName, e.folderSelect)

	e.Container = container.NewBorder(topBar, nil, nil, nil, treeFiles)

	return e

}
