package menu

import (
	"fyne.io/fyne"
)

func createFileMenu() (menu *fyne.Menu) {

	openfolder := fyne.NewMenuItem("Open Folder", func() {})
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

func NewMenuBar() (menubar *fyne.MainMenu) {

	fileMenu := createFileMenu()
	editMenu := createEditMenu()

	menubar = fyne.NewMainMenu(fileMenu, editMenu)

	return
}
