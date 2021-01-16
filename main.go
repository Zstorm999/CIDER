package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"github.com/Zstorm999/cider/explorer"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("CIDER")

	explorer := explorer.New(myWindow)

	myWindow.SetContent(explorer)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
