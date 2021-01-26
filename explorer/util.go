package explorer

import (
	"errors"
	"os"
	"strings"
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

func parseFileName(file string) string {
	if file == "" {
		return ""
	}

	if file == "/" || file == "file:///" {
		return "/"
	}

	sliced := strings.Split(file, "/")
	n := len(sliced)

	return sliced[n-1]
}
