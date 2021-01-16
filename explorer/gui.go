package explorer

import (
	"fmt"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
)

//New : returns a fyne.Container containing a working file explorer
func New(win fyne.Window) *fyne.Container {

	path, err := os.Getwd()

	if err != nil {
		fmt.Println(err)
		path = ""
	}

	lbcurrentFolder := createFolderWidget(path, win)

	treeFiles := createFilesList(path)

	container := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		lbcurrentFolder, treeFiles)

	return container
}
