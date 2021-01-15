package explorer

import (
	"errors"
	"fmt"
	"os"

	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
)

func getDirectoryContent(path string) ([]os.FileInfo, error) {

	currentDir, err := os.Open(path)
	defer currentDir.Close()

	if err != nil {
		return nil, err
	}

	//checking that the opened file is a directory
	dirStats, err := currentDir.Stat()
	if err != nil {
		return nil, err
	} else if !dirStats.IsDir() {
		return nil, errors.New("Specified path is not a directory: how could that ever happen ?")
	}

	//putting all files in the directory in a slice
	files, err := currentDir.Readdir(0)

	if err != nil {
		return nil, err
	}

	return files, nil

}

func sortFileList(list []os.FileInfo) []os.FileInfo {

	var dirList, otherList []os.FileInfo

	for i := range list {
		if list[i].IsDir() {
			dirList = append(dirList, list[i])
		} else {
			otherList = append(otherList, list[i])
		}
	}
	return append(dirList, otherList...)

}

func createFilesList(path string) *widget.List {

	files, err := getDirectoryContent(path)

	files = sortFileList(files)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	list := widget.NewList(
		func() int { return len(files) }, //returns the length of the list
		func() fyne.CanvasObject { //returns a template element, use as default
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(nil), widget.NewLabel("Template"))
		},
		func(i widget.ListItemID, o fyne.CanvasObject) { //creates an element from the template, given its ID in the list
			o.(*fyne.Container).Objects[1].(*widget.Label).SetText(files[i].Name())

			if files[i].IsDir() {
				o.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.FolderIcon())
			} else {
				o.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.DocumentIcon())
			}

		},
	)

	return list

}
