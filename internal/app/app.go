package app

import (
	"github.com/cilium/ebpf"
	"github.com/rivo/tview"
)

type Application struct {
	CurrentView       uint32
	SelectedProgramID ebpf.ProgramID
	App               *tview.Application
	MapsView          *tview.Flex
	ProgListView      *tview.Flex
}

func (a *Application) SetCurrentView(handler func()) {
	handler()
}
func (a *Application) GetCurrentView() uint32 {
	return a.CurrentView
}
func (a *Application) SetSelectedProgramID(id ebpf.ProgramID) {
	a.SelectedProgramID = id
}
func (a *Application) NewApplication() *tview.Application {
	app := tview.NewApplication()
	app.EnableMouse(true)

	a.App = app
	return app
}
