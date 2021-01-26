package explorer

import (
	"fmt"
	"strings"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func extendFilesMap(path string, files *map[string][]string, completion *widget.ProgressBar, level float64) int {

	filesRaw, err := getDirectoryContent(path)

	if err != nil {
		fmt.Println("Error while loading files from memory")
	}

	filesRaw = sortFileList(filesRaw)

	nbFiles := len(filesRaw)

	nextLevel := level / float64(nbFiles)

	for i := range filesRaw {

		if filesRaw[i].Name() != ".git" {
			var file string

			if filesRaw[i].IsDir() {
				nbNext := extendFilesMap(path+"/"+filesRaw[i].Name(), files, completion, nextLevel)

				file = "dir:" + path + "/" + filesRaw[i].Name()

				if nbNext != 0 {
					completion.SetValue(completion.Value + nextLevel)
				}

			} else {
				file = "doc:" + path + "/" + filesRaw[i].Name()

				completion.SetValue(completion.Value + nextLevel)
			}

			(*files)["dir:"+path] = append((*files)["dir:"+path], file)

		}

	}

	return nbFiles

}

func createFilesMap(path string, completion *widget.ProgressBar) (files map[string][]string) {

	files = make(map[string][]string)

	completion.SetValue(0)

	extendFilesMap(path, &files, completion, 1)

	//placing the root
	files[""] = append(files[""], "dir:"+path)

	return

}

func createFilesTree(path string, completion *widget.ProgressBar) *widget.Tree {

	files := createFilesMap(path, completion)

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

			if strings.HasPrefix(uid, "dir:") {
				node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.FolderIcon())
			} else {
				node.(*fyne.Container).Objects[0].(*widget.Icon).SetResource(theme.DocumentIcon())
			}

			node.(*fyne.Container).Objects[1].(*widget.Label).SetText(parseFileName(uid))
		})

	return tree

}
