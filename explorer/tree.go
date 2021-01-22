package explorer

import (
	"fmt"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func extendFilesMap(path string, files *map[string][]string) {

	filesRaw, err := getDirectoryContent(path)

	if err != nil {
		fmt.Println("Error while loading files from memory")
	}

	filesRaw = sortFileList(filesRaw)

	for i := range filesRaw {

		if filesRaw[i].Name() != ".git" {

			(*files)[path] = append((*files)[path], path+"/"+filesRaw[i].Name())

			if filesRaw[i].IsDir() {
				extendFilesMap(path+"/"+filesRaw[i].Name(), files)
			}
		}

	}

}

func createFilesMap(path string) (files map[string][]string) {

	files = make(map[string][]string)
	extendFilesMap(path, &files)

	//placing the root
	files[""] = append(files[""], path)

	return

}

func createFilesTree(path string) *widget.Tree {

	files := createFilesMap(path)

	tree := widget.NewTree(
		func(uid string) (children []string) { //gets all children from a node, using its UID
			for _, elt := range files[uid] {
				children = append(children, elt)
			}
			return
		},
		func(uid string) bool { //returns true if the node is a leaf
			_, exist := files[uid]
			return exist
		},
		func(branch bool) fyne.CanvasObject { //creates a template object
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(nil), widget.NewLabel(""))
		},
		func(uid string, branch bool, node fyne.CanvasObject) { //update the template object to create what will be displayed
			if uid == path {
				node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.FolderIcon())
				node.(*fyne.Container).Objects[1].(*widget.Label).SetText(uid)
			} else {
				sliced := strings.Split(uid, "/")
				n := len(sliced)

				_, isFolder := files[uid]

				if isFolder {
					node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.FolderIcon())
				} else {
					node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.DocumentIcon())
				}

				node.(*fyne.Container).Objects[1].(*widget.Label).SetText(sliced[n-1])
			}
		})

	return tree

	//return createFilesList(path)

}
