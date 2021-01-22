package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"github.com/Zstorm999/cider/explorer"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("CIDER")

	explorer := explorer.New(myWindow)
	lbHello := widget.NewLabel("Hello World")

	windowContent := fyne.NewContainerWithLayout(layout.NewHBoxLayout(), explorer, lbHello)

	myWindow.SetContent(windowContent)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
