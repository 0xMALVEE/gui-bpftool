package main

import (
	"tui-bpftool/internal/proglist"

	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	app.EnableMouse(true)

	progListView := proglist.GetProgListView(app)

	if err := app.SetRoot(progListView, true).Run(); err != nil {
		panic(err)
	}
}
