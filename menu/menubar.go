package menu

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"github.com/Zstorm999/cider/explorer"
)

func createFileMenu(win fyne.Window, e *explorer.Explorer) (menu *fyne.Menu) {

	openfolder := fyne.NewMenuItem("Open Folder",
		func() {
			selectWindow := dialog.NewFolderOpen(
				func(file fyne.ListableURI, err error) {
					e.UpdateTree(file.String())
				}, win)

			selectWindow.Show()
		})

	openfile := fyne.NewMenuItem("Open File", func() {})

	menu = fyne.NewMenu("File", openfile, openfolder)

	return
}

func createEditMenu() (menu *fyne.Menu) {

	undo := fyne.NewMenuItem("Undo", func() {})
	redo := fyne.NewMenuItem("Redo", func() {})

	menu = fyne.NewMenu("Edit", undo, redo)

	return
}

func NewMenuBar(win fyne.Window, e *explorer.Explorer) (menubar *fyne.MainMenu) {

	fileMenu := createFileMenu(win, e)
	editMenu := createEditMenu()

	menubar = fyne.NewMainMenu(fileMenu, editMenu)

	return
}
