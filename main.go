package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"github.com/Zstorm999/cider/explorer"
	"github.com/Zstorm999/cider/menu"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("CIDER")

	explorer := explorer.New(myWindow)
	lbHello := widget.NewLabel("Hello World")

	mainFrame := container.NewHSplit(explorer.Container, lbHello)

	myWindow.SetContent(mainFrame)

	menubar := menu.NewMenuBar()
	myWindow.SetMainMenu(menubar)

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
