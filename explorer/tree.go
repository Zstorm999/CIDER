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
			var file string

			if filesRaw[i].IsDir() {
				extendFilesMap(path+"/"+filesRaw[i].Name(), files)

				file = "dir:" + path + "/" + filesRaw[i].Name()
			} else {
				file = "doc:" + path + "/" + filesRaw[i].Name()
			}

			(*files)["dir:"+path] = append((*files)["dir:"+path], file)
		}

	}

}

func createFilesMap(path string) (files map[string][]string) {

	files = make(map[string][]string)
	extendFilesMap(path, &files)

	//placing the root
	files[""] = append(files[""], "dir:"+path)

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
		func(uid string) bool { //returns false if the node is a leaf

			if uid == "" {
				return true
			}

			return strings.HasPrefix(uid, "dir:")
		},
		func(branch bool) fyne.CanvasObject { //creates a template object
			return fyne.NewContainerWithLayout(layout.NewHBoxLayout(), widget.NewIcon(nil), widget.NewLabel(""))
		},
		func(uid string, branch bool, node fyne.CanvasObject) { //update the template object to create what will be displayed

			sliced := strings.Split(uid, "/")
			n := len(sliced)

			if strings.HasPrefix(uid, "dir:") {
				node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.FolderIcon())
			} else {
				node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.DocumentIcon())
			}

			node.(*fyne.Container).Objects[1].(*widget.Label).SetText(sliced[n-1])
		})

	return tree

}
