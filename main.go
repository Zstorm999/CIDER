package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/Zstorm999/cider/editor"
	"github.com/Zstorm999/cider/explorer"
	"github.com/Zstorm999/cider/menu"
	"github.com/Zstorm999/cider/resources"
)

func main() {

	myApp := app.New()
	myApp.SetIcon(resources.CiderIcon())

	myWindow := myApp.NewWindow("CIDER")

	explorer := explorer.New(myWindow)

	editor := editor.New()

	mainFrame := container.NewHSplit(explorer.Container, editor.Container)
	mainFrame.SetOffset(0.2)

	myWindow.SetContent(mainFrame)

	menubar := menu.NewMenuBar(myWindow, explorer, editor)
	myWindow.SetMainMenu(menubar)

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
