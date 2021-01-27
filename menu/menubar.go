package menu

import (
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"github.com/Zstorm999/cider/editor"
	"github.com/Zstorm999/cider/explorer"
)

func createFileMenu(win fyne.Window, explorer *explorer.Explorer, editor *editor.Editor) (menu *fyne.Menu) {

	openfolder := fyne.NewMenuItem("Open Folder",
		func() {
			selectWindow := dialog.NewFolderOpen(
				func(file fyne.ListableURI, err error) {
					if file != nil {
						explorer.UpdateTree(file.String())
					}
				}, win)

			selectWindow.Show()
		})

	openfile := fyne.NewMenuItem("Open File",
		func() {
			selectWindow := dialog.NewFileOpen(
				func(file fyne.URIReadCloser, err error) {

					if file != nil {
						filepath := strings.TrimPrefix(file.URI().String(), "file://")
						editor.AddTab(filepath)
					}

				}, win)
			selectWindow.Show()
		})

	menu = fyne.NewMenu("File", openfile, openfolder)

	return
}

func createEditMenu() (menu *fyne.Menu) {

	undo := fyne.NewMenuItem("Undo", func() {})
	redo := fyne.NewMenuItem("Redo", func() {})

	menu = fyne.NewMenu("Edit", undo, redo)

	return
}

func NewMenuBar(win fyne.Window, explorer *explorer.Explorer, editor *editor.Editor) (menubar *fyne.MainMenu) {

	fileMenu := createFileMenu(win, explorer, editor)
	editMenu := createEditMenu()

	menubar = fyne.NewMainMenu(fileMenu, editMenu)

	return
}
