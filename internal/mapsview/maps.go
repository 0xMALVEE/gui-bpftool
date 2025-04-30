package mapsview

import (
	"strconv"
	"tui-bpftool/internal"
	"tui-bpftool/internal/app"
	bpfmaps "tui-bpftool/pkg/bpf-maps"

	"github.com/rivo/tview"
)

func GetMapsView(app *app.Application) *tview.Flex {
	programId := app.SelectedProgramID
	button := tview.NewButton("Back").SetSelectedFunc(func() {
		app.SetCurrentView(func() {
			app.CurrentView = internal.ProgramListView
			app.App.SetRoot(app.ProgListView, true)
		})
	})

	maps, _ := bpfmaps.GetAllMaps(programId)

	length := strconv.Itoa(len(maps))

	textView := tview.NewTextArea().SetText(length, false)

	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(button, 1, 0, false).AddItem(textView, 0, 1, false)

	return layout
}
