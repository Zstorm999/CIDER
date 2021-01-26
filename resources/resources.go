package resources

import (
	"fmt"

	"fyne.io/fyne"
)

func CiderIcon() fyne.Resource {

	icon, err := fyne.LoadResourceFromPath("resources/icons/ciderglass.png")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return icon
}
