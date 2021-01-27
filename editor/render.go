package editor

import (
	"image/color"

	"fyne.io/fyne/v2"
)

var (
	backgroundColor = color.RGBA{R: 30, G: 30, B: 30, A: 255}
)

type render struct {
	editor *editorTab
}

func (r *render) Layout(s fyne.Size) {
	//r.editor.content.Resize(s)
}

func (r *render) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (r *render) Refresh() {
	//currently does nothing
}

func (r *render) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.editor.content}
}

func (r *render) BackgroundColor() color.Color {
	return backgroundColor
}

func (r *render) Destroy() {
	//currently does nothing
}

func (e *editorTab) CreateRenderer() fyne.WidgetRenderer {

	r := &render{e}

	return r
}
