package main

import (
	bpfprog "gui-bpftool/pkg/bpf-prog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("guibpftool")
	w.Resize(fyne.NewSize(800, 600))

	ebpgProgList, _ := bpfprog.GetProgListWithInfo()

	list := widget.NewList(
		func() int {
			return len(ebpgProgList)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(ebpgProgList[i].ProgramInfo.Name)
		})

	w.SetContent(list)
	w.ShowAndRun()
}
