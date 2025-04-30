package main

import (
	bpfprog "tui-bpftool/pkg/bpf-prog"

	"github.com/rivo/tview"
)

func main() {

	// box := tview.NewBox().SetBorder(true).SetTitle("tui-bpftool!")

	table := tview.NewTable().SetBorders(true).SetSelectable(true, false)

	table.SetCell(0, 0, tview.NewTableCell("Programs List").SetTextColor(tview.Styles.PrimaryTextColor))
	table.SetCell(0, 1, tview.NewTableCell("Type").SetTextColor(tview.Styles.PrimaryTextColor))

	ebpgProgList, _ := bpfprog.GetProgListWithInfo()

	for i, prog := range ebpgProgList {
		table.SetCell(i+1, 0, tview.NewTableCell(prog.ProgramInfo.Name).SetTextColor(tview.Styles.PrimaryTextColor))
		table.SetCell(i+1, 1, tview.NewTableCell(prog.Type).SetTextColor(tview.Styles.PrimaryTextColor))
	}

	if err := tview.NewApplication().SetRoot(table, true).Run(); err != nil {
		panic(err)
	}
}
