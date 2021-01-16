package explorer

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func createFolderWidget(path string, win fyne.Window) *fyne.Container {

	name := widget.NewLabel(path)

	selectWindow := dialog.NewFolderOpen(func(file fyne.ListableURI, err error) {
		name.SetText(file.Name())
	}, win)

	btSelect := widget.NewButtonWithIcon("", theme.FolderIcon(),
		func() {
			selectWindow.Show()
		})

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), name, btSelect)

}
