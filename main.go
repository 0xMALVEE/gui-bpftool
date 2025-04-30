package main

import (
	bpfprog "tui-bpftool/pkg/bpf-prog"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	searchBpf := tview.NewInputField().
		SetLabel("Search: ").
		SetPlaceholder("Enter program name").
		SetFieldWidth(30)

	bpfListTable := tview.NewTable().SetSelectable(true, false)

	fillBpfListTable := func(programName string) {
		bpfListTable.Clear()
		bpfListTable.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(tview.Styles.PrimaryTextColor))
		bpfListTable.SetCell(0, 1, tview.NewTableCell("Type").SetTextColor(tview.Styles.PrimaryTextColor))

		ebpgProgList, _ := bpfprog.GetProgListWithInfo(programName)

		for i, prog := range ebpgProgList {
			bpfListTable.SetCell(i+1, 0, tview.NewTableCell(prog.ProgramInfo.Name).SetTextColor(tview.Styles.PrimaryTextColor))
			bpfListTable.SetCell(i+1, 1, tview.NewTableCell(prog.Type).SetTextColor(tview.Styles.PrimaryTextColor))
		}
	}

	fillBpfListTable("")

	searchBpf.SetChangedFunc(func(text string) {
		if text == "" {
			fillBpfListTable("")
			return
		}
		fillBpfListTable(text)
	})

	layout := tview.NewFlex().SetDirection(tview.FlexRow).AddItem(searchBpf, 1, 0, false).AddItem(bpfListTable, 0, 1, true)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyTab {
			if bpfListTable.HasFocus() {
				app.SetFocus(searchBpf)
				return nil
			}
			app.SetFocus(bpfListTable)
			return nil
		}
		return event
	})

	if err := app.SetRoot(layout, true).SetFocus(searchBpf).Run(); err != nil {
		panic(err)
	}
}
