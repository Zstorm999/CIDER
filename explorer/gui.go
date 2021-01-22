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
	//path := "/home/thomas"

	//currentFolder := fyne.NewContainerWithoutLayout()
	container := fyne.NewContainerWithoutLayout()

	//currentFolder := createFolderWidget(path, win, func(s string) { container.Objects[1] = createFilesTree(s) })
	//container.Add(currentFolder)

	treeFiles := createFilesTree(path)
	container.Add(treeFiles)

	container.Layout = layout.NewGridLayout(1)

	return container
}
