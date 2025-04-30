package mapsview

import (
	"tui-bpftool/internal"
	"tui-bpftool/internal/app"

	"github.com/rivo/tview"
)

func GetMapsView(app *app.Application) *tview.Flex {
	button := tview.NewButton("Back").SetSelectedFunc(func() {
		app.SetCurrentView(func() {
			app.CurrentView = internal.ProgramListView
			app.App.SetRoot(app.ProgListView, true)
		})
	})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(button, 1, 0, false)

	return layout
}
