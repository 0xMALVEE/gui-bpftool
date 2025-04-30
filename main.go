package main

import (
	bpfprog "tui-bpftool/pkg/bpf-prog"

	"github.com/rivo/tview"
)

func main() {
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

	if err := tview.NewApplication().SetRoot(layout, true).SetFocus(searchBpf).Run(); err != nil {
		panic(err)
	}
}
