package main

import (
	"tui-bpftool/internal"
	"tui-bpftool/internal/app"
	"tui-bpftool/internal/mapsview"
	"tui-bpftool/internal/proglist"
)

func main() {
	app := app.Application{CurrentView: internal.ProgramListView}
	app.NewApplication()

	// set the views
	app.ProgListView = proglist.GetProgListView(&app)
	app.MapsView = mapsview.GetMapsView(&app)

	if err := app.App.SetRoot(app.ProgListView, true).Run(); err != nil {
		panic(err)
	}
}
