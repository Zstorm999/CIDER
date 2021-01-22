package explorer

import (
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func createFolderWidget(path string, win fyne.Window, updateTree func(string)) *fyne.Container {

	name := widget.NewLabel(path)

	selectWindow := dialog.NewFolderOpen(func(file fyne.ListableURI, err error) {
		trimmed := strings.TrimPrefix(file.String(), "file://")
		name.SetText(trimmed)

		updateTree(trimmed)

	}, win)

	btSelect := widget.NewButtonWithIcon("", theme.FolderIcon(),
		func() {
			selectWindow.Show()
		})

	return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), name, btSelect)

}
