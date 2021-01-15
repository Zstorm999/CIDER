package explorer

import (
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

//New : returns a fyne.Container containing a working file explorer
func New() *fyne.Container {

	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		path = ""
	}

	lbcurrentFolder := widget.NewLabel(path)

	treeFiles := createFilesList(path)

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(1),
		lbcurrentFolder, treeFiles)

	return container
}
